// Package route encapsulates HTTP routing and handlers.
package route

import (
	"context"
	"net/http"

	"github.com/motki/motki-server/http/auth"
	"github.com/motki/motki-server/http/session"
	"github.com/motki/motki/log"
)

// A Request contains information about an incoming request.
type Request struct {
	*http.Request

	session *session.Session
	auth    *auth.Session
	authCtx context.Context
}

// Session returns the session associated with the request.
func (r *Request) Session() (*session.Session, bool) {
	return r.session, r.session != nil
}

// Auth returns the authenticated session associated with the request.
func (r *Request) Auth() (*auth.Session, bool) {
	return r.auth, r.auth != nil
}

// AuthorizedContext returns an appropriate authorized context for the request.
func (r *Request) AuthorizedContext() (context.Context, bool) {
	return r.authCtx, r.authCtx != nil
}

// SetSession sets the request's associated session.
func (r *Request) SetSession(sess *session.Session) {
	r.session = sess
}

// SetAuth sets the request's associated authenticated session.
func (r *Request) SetAuth(a *auth.Session) {
	r.auth = a
}

// SetAuthorizedContext sets the request's authorized Context for use in other services.
func (r *Request) SetAuthorizedContext(ctx context.Context) {
	r.authCtx = ctx
}

// ServeMux is a wrapper around a stdlib http.ServeMux.
//
// Handlers for a ServeMux receive a specialized Request defined in this package.
// Middlewares can be used to handle, for example, session population or authorization.
type ServeMux struct {
	*http.ServeMux

	logger log.Logger
}

// NewServeMux creates a new ServeMux, ready for use.
//
// Order of the middlewares is significant. First provided will execute first,
// and so on.
func NewServeMux(l log.Logger) *ServeMux {
	return &ServeMux{
		ServeMux: http.NewServeMux(),
		logger:   l,
	}
}

// Handle registers a handler for the given pattern.
func (mux *ServeMux) Handle(pattern string, h Handler) {
	mux.ServeMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		req := &Request{
			Request: r,
		}
		if err := h.ServeHTTP(w, req); err != nil {
			mux.logger.Warnf("handler returned error: %s", err.Error())
		}
	})
}

// HandleFunc is a convenience method for registering a closure handler.
func (mux *ServeMux) HandleFunc(pattern string, h func(w http.ResponseWriter, req *Request) error) {
	mux.Handle(pattern, HandlerFunc(h))
}

// A Handler responds to a given HTTP request.
type Handler interface {
	// ServeHTTP handles the given request.
	ServeHTTP(w http.ResponseWriter, req *Request) error
}

// HandlerFunc is a closure that can be used as a Handler
type HandlerFunc func(w http.ResponseWriter, req *Request) error

func (r HandlerFunc) ServeHTTP(w http.ResponseWriter, req *Request) error {
	return r(w, req)
}
