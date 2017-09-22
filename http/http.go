// Package http manages the motki application web server.
package http

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/motki/motki/log"
	"github.com/motki/motki-server/http/route"
	"github.com/motki/motki-server/http/session"
	"github.com/motki/motki-server/http/template"
	"golang.org/x/crypto/acme/autocert"
)

// A module represents a contained set of http Handlers.
//
// Modules should manage their own dependencies and should never
// have dependencies on other Modules.
type module interface {
	// Init allows the module to bind any handlers to the given ServeMux.
	Init(mux *route.ServeMux) error
}

// Config defines the basic configuration for the web server.
type Config struct {
	SSL        SSLConfig       `toml:"ssl"`
	Session    session.Config  `toml:"sessions"`
	Templating template.Config `toml:"templates"`

	ListenAddr string `toml:"listen"`
	Redirect   bool   `toml:"redirect"`

	listenHost string
	listenPort string
}

// SSLConfig defines a TLS configuration for the web server.
type SSLConfig struct {
	ListenAddr string   `toml:"listen"`
	AutoCert   bool     `toml:"autocert"`
	CertFile   string   `toml:"certfile"`
	CertKey    string   `toml:"keyfile"`
	RequireSSL bool     `toml:"require"`
	ExtraHosts []string `toml:"extra_hosts"`

	listenHost string
	listenPort string
}

// TLSConfig attempts to load the configured certificate.
func (c SSLConfig) TLSConfig() (*tls.Config, error) {
	if c.AutoCert {
		hosts := append([]string{c.listenHost}, c.ExtraHosts...)
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(hosts...),
			Cache:      autocert.DirCache("certs"),
		}
		return &tls.Config{GetCertificate: m.GetCertificate}, nil
	}
	var err error
	tc := &tls.Config{}
	tc.NextProtos = []string{"http", "h2"}
	tc.Certificates = make([]tls.Certificate, 1)
	tc.Certificates[0], err = tls.LoadX509KeyPair(c.CertFile, c.CertKey)
	if err != nil {
		return nil, err
	}
	return tc, nil
}

// Server is a wrapper used to orchestrate functions of the web server.
type Server struct {
	conf   Config
	logger log.Logger

	mux *route.ServeMux

	web *http.Server
	ssl *http.Server
}

// New creates a new web server instance with the given configuration.
func New(c Config, l log.Logger) (*Server, error) {
	var err error
	c.listenHost, c.listenPort, err = net.SplitHostPort(c.ListenAddr)
	if err != nil {
		return nil, err
	}
	c.SSL.listenHost, c.SSL.listenPort, err = net.SplitHostPort(c.SSL.ListenAddr)
	if err != nil {
		return nil, err
	}
	mux := route.NewServeMux(l)
	return &Server{conf: c, logger: l, mux: mux}, nil
}

// Register initializes the given modules.
func (s *Server) Register(mods ...module) error {
	for _, m := range mods {
		err := m.Init(s.mux)
		if err != nil {
			return err
		}
	}
	return nil
}

// ListenAndServe begins listening for HTTP requests.
func (s *Server) ListenAndServe() error {
	if s.web != nil {
		return fmt.Errorf("http server already started")
	}
	var mux http.Handler = s.mux
	var listenAddr string = s.conf.ListenAddr
	if s.conf.SSL.RequireSSL {
		// A static redirect handler that redirects the user to the SSL site.
		s.logger.Debugf("http: ssl required; all http requests will be forwarded to https")
		mux = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.URL.Scheme = "https"
			r.URL.Host = s.conf.SSL.ListenAddr
			http.Redirect(w, r, r.URL.String(), http.StatusPermanentRedirect)
		})
	} else if s.conf.Redirect {
		// A special handler to ensure the user hits the canonical hostname.
		s.logger.Debugf("http: listening for any hostname on port: %s", s.conf.listenPort)
		s.logger.Debugf("http: canonical hostname redirection enabled, canonical host: %s", s.conf.SSL.listenHost)
		listenAddr = ":" + s.conf.listenPort
		mux = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if stripPort(r.Host) != s.conf.listenHost {
				r.URL.Host = s.conf.ListenAddr
				http.Redirect(w, r, r.URL.String(), http.StatusPermanentRedirect)
				return
			}
			s.mux.ServeHTTP(w, r)
		})
	} else {
		s.logger.Debugf("http: listening on: %s", s.conf.ListenAddr)
	}
	logger, closer, err := log.StdLogger(s.logger, "warn")
	if err != nil {
		s.logger.Warnf("http: unable to create stdlib Logger: %s", err.Error())
	}
	if closer != nil {
		// TODO: Close returns an error
		defer closer.Close()
	}
	s.web = &http.Server{
		Addr:     listenAddr,
		Handler:  mux,
		ErrorLog: logger,
	}
	defer func() {
		// Clean up.
		s.web = nil
	}()
	return s.web.ListenAndServe()
}

// ListenAndServeTLS begins listening for HTTPS requests.
func (s *Server) ListenAndServeTLS() error {
	if s.ssl != nil {
		return fmt.Errorf("https server already started")
	}
	tc, err := s.conf.SSL.TLSConfig()
	if err != nil {
		return err
	}
	var mux http.Handler = s.mux
	var listenAddr string = s.conf.SSL.ListenAddr
	if s.conf.Redirect {
		// A special handler to ensure the user hits the canonical hostname.
		s.logger.Debugf("https: listening for any hostname on port: %s", s.conf.SSL.listenPort)
		s.logger.Debugf("https: canonical hostname redirection enabled, canonical host: %s", s.conf.SSL.listenHost)
		listenAddr = ":" + s.conf.SSL.listenPort
		mux = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if stripPort(r.Host) != s.conf.SSL.listenHost {
				r.URL.Host = s.conf.SSL.ListenAddr
				http.Redirect(w, r, r.URL.String(), http.StatusPermanentRedirect)
				return
			}
			s.mux.ServeHTTP(w, r)
		})
	} else {
		s.logger.Debugf("https: listening on: %s", s.conf.SSL.ListenAddr)
	}
	logger, closer, err := log.StdLogger(s.logger, "warn")
	if err != nil {
		s.logger.Warnf("https: unable to create stdlib Logger: %s", err.Error())
	}
	if closer != nil {
		// TODO: Close returns an error
		defer closer.Close()
	}
	if s.conf.SSL.AutoCert {
		s.logger.Debugf("https: using acme/autocert to generate ssl certificates")
		s.logger.Debugf("https: autocert domain whitelist: %v", s.conf.SSL.ExtraHosts)
	} else {
		s.logger.Debugf("https: ssl cert file: %s", s.conf.SSL.CertFile)
		s.logger.Debugf("https: ssl key file: %s", s.conf.SSL.CertKey)
	}
	s.ssl = &http.Server{
		Addr:      listenAddr,
		Handler:   mux,
		ErrorLog:  logger,
		TLSConfig: tc,
	}
	defer func() {
		// Clean up.
		s.ssl = nil
	}()
	return s.ssl.ListenAndServeTLS(s.conf.SSL.CertFile, s.conf.SSL.CertKey)
}

// Shutdown attempts to shutdown both the HTTP and HTTPS servers gracefully.
func (s *Server) Shutdown() error {
	timeout := 2 * time.Second
	var werr error
	var serr error
	if s.web != nil {
		s.logger.Debugf("http: shutting down with %.2f sec timeout", timeout.Seconds())
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		werr = s.web.Shutdown(ctx)
	}
	if s.ssl != nil {
		s.logger.Debugf("https: shutting down with %.2f sec timeout", timeout.Seconds())
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		serr = s.ssl.Shutdown(ctx)
	}
	if werr != nil || serr != nil {
		return fmt.Errorf("http returned error: %v; https returned error: %v", werr, serr)
	}
	return nil
}

// From: https://github.com/golang/go/blob/d7ec89c19846d8c1d89d510cd7634ae9de640ac0/src/net/url/url.go#L1014
func stripPort(hostport string) string {
	colon := strings.IndexByte(hostport, ':')
	if colon == -1 {
		return hostport
	}
	if i := strings.IndexByte(hostport, ']'); i != -1 {
		return strings.TrimPrefix(hostport[:i], "[")
	}
	return hostport[:colon]
}
