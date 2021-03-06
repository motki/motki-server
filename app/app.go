// Package app contains functionality related to creating an integrated
// motki-server environment with all the necessary dependencies.
//
// The goal with this package is to provide a single, reusable base for
// getting a motki application server up and running.
//
// This package imports every other motki-server package. As such, it cannot
// be imported from the "library" portion of the project. It is intended to be
// used from an external package (as is done in the motkid command).
//
// This package provides a web application server and can optionally serve
// as a remote backend for client-only motki applications.
package app

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"

	"github.com/motki/core/app"
	"github.com/motki/motki-server/http"
	"github.com/motki/motki-server/http/auth"
	_ "github.com/motki/motki-server/http/middleware"
	_ "github.com/motki/motki-server/http/route"
	"github.com/motki/motki-server/http/session"
	"github.com/motki/motki-server/http/template"
	"github.com/motki/motki-server/mail"

	modaccount "github.com/motki/motki-server/http/module/account"
	modassets "github.com/motki/motki-server/http/module/assets"
	modauth "github.com/motki/motki-server/http/module/auth"
	modhome "github.com/motki/motki-server/http/module/home"
	modindustry "github.com/motki/motki-server/http/module/industry"
	modmarket "github.com/motki/motki-server/http/module/market"
)

// A WebEnv wraps a regular Env, providing web and mail servers.
type WebEnv struct {
	*app.Env

	Mailer    *mail.Sender
	Sessions  session.Manager
	Templates template.Renderer
	Auth      auth.Manager
	Web       *http.Server

	signals chan os.Signal
}

// Config represents the configuration for a motki application server.
type Config struct {
	*app.Config

	HTTP http.Config `toml:"http"`
	Mail mail.Config `toml:"mail"`
}

// NewConfigFromTOMLFile loads a TOML configuration from the given path.
func NewConfigFromTOMLFile(tomlPath string) (*Config, error) {
	if !filepath.IsAbs(tomlPath) {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		tomlPath = filepath.Join(cwd, tomlPath)
	}
	c, err := ioutil.ReadFile(tomlPath)
	if err != nil {
		return nil, err
	}
	conf := &Config{}
	_, err = toml.Decode(string(c), conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// NewWebEnv creates a new web environment using the given configuration.
//
// This function will initialize a regular Env before it initializes the
// web and mail server related functionality.
func NewWebEnv(conf *Config) (*WebEnv, error) {
	env, err := app.NewEnv(conf.Config)
	if err != nil {
		return nil, err
	}
	mailer := mail.NewSender(conf.Mail, env.Logger)
	mailer.DoNotSend, err = mail.NewModelList(env.Model, "unsubscribe")
	if err != nil {
		return nil, errors.Wrap(err, "app: unable to init 'unsubscribe' list")
	}
	sessions := session.NewManager(conf.HTTP.Session, env.Logger)
	templates := template.NewRenderer(conf.HTTP.Templating, env.Logger)
	authManager := auth.NewManager(
		auth.NewFormLoginAuthenticator(env.Model, env.Logger, "/login/begin"),
		auth.NewEveAPIAuthorizer(env.Model, env.EveAPI, env.Logger),
		sessions)
	srv, err := http.New(conf.HTTP, env.Logger)
	if err != nil {
		return nil, errors.Wrap(err, "app: unable to initialize web environment")
	}

	err = srv.Register(
		modassets.New(),
		modauth.New(sessions, authManager, templates, env.Model, env.Scheduler, mailer, env.Logger),
		modhome.New(sessions, templates, mailer, env.Logger),
		modmarket.New(authManager, templates, env.Client, env.Logger),
		modaccount.New(authManager, templates, env.Model, env.EveDB, env.Scheduler, env.Logger),
		modindustry.New(authManager, templates, env.Model, env.EveDB, env.Logger))
	if err != nil {
		return nil, errors.Wrap(err, "app: unable to initialize web environment")
	}
	return &WebEnv{
		Env: env,

		Mailer:    mailer,
		Sessions:  sessions,
		Templates: templates,
		Auth:      authManager,
		Web:       srv,
	}, nil
}

// BlockUntilSignal will block until it receives the signals signal.
//
// This function performs the default shutdown procedure when it receives
// an signals signal.
//
// See BlockUntilSignalWith on Env for more details.
func (webEnv *WebEnv) BlockUntilSignal(abort chan os.Signal) {
	webEnv.signals = abort
	shutdownFuncs := []app.ShutdownFunc{
		func() {
			if err := webEnv.Scheduler.Shutdown(); err != nil {
				webEnv.Logger.Warnf("app: error shutting down scheduler: %s", err.Error())
			}
		},
		func() {
			if webEnv.Server == nil {
				return
			}
			if err := webEnv.Server.Shutdown(); err != nil {
				webEnv.Logger.Warnf("app: error shutting down grpc server: %s", err.Error())
			}
		},
		func() {
			if err := webEnv.Web.Shutdown(); err != nil {
				webEnv.Logger.Warnf("app: error shutting down web server: %s", err.Error())
			}
		}}
	webEnv.BlockUntilSignalWith(abort, shutdownFuncs...)
}
