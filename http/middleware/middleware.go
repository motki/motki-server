// Package middleware contains HTTP middlewares for Session and Auth.
package middleware

import (
	"errors"
	"net/http"

	"github.com/motki/core/model"
	"github.com/motki/motki-server/http/auth"
	"github.com/motki/motki-server/http/route"
	"github.com/motki/motki-server/http/session"
)

// Sessionize wraps the given handler in a session middleware.
//
// This middleware is responsible for populating a request's session.
func Sessionize(s session.Manager, h route.Handler) route.Handler {
	return route.HandlerFunc(func(w http.ResponseWriter, req *route.Request) (err error) {
		sess, err := s.Get(req.Request, w)
		if err != nil {
			// TODO: poor way to handle the error, shouldn't occur in practice.
			w.WriteHeader(500)
			return err
		}
		defer func() {
			if ferr := sess.Flush(); ferr != nil && err == nil {
				err = ferr
			}
		}()
		req.SetSession(sess)
		err = h.ServeHTTP(w, req)
		return err
	})
}

// SessionizeFunc is a convenience function for wrapping a closure in a sessionize middleware.
func SessionizeFunc(s session.Manager, h func(w http.ResponseWriter, req *route.Request) error) route.Handler {
	return Sessionize(s, route.HandlerFunc(h))
}

// Authenticate wraps the given handler in an authentication middleware.
//
// This middleware ensures the user has logged into a user account before they
// can access the given handler.
//
// The Authenticate middleware implicitly applies the Sessionize middleware.
func Authenticate(m auth.Manager, h route.Handler) route.Handler {
	return Authorize(m, model.RoleAnon, h)
}

// AuthenticateFunc is a convenience function for wrapping a closure
// in an authentication middleware.
//
// The Authenticate middleware implicitly applies the Sessionize middleware.
func AuthenticateFunc(m auth.Manager, h func(w http.ResponseWriter, req *route.Request) error) route.Handler {
	return Authenticate(m, route.HandlerFunc(h))
}

// Authorize wraps the given handler in an authorization middleware.
//
// This middleware ensures the user has the given role necessary to access
// the given handler.
//
// The Authorize middleware implicitly applies both Authenticate and Sessionize middlewares.
func Authorize(m auth.Manager, r model.Role, h route.Handler) route.Handler {
	return Sessionize(m.Sessions, route.HandlerFunc(func(w http.ResponseWriter, req *route.Request) error {
		sess, ok := req.Session()
		if !ok {
			// TODO: poor way to handle the error, shouldn't occur in practice.
			w.WriteHeader(500)
			return errors.New("expected session")
		}
		a, ok := m.VerifyAuthentication(sess)
		if !ok {
			m.BeginAuthentication(sess, w, req.Request)
			return nil
		}
		req.SetAuth(a)
		if r != model.RoleAnon {
			ctx, err := m.AuthorizedContext(a, r)
			if err != nil {
				sess.Set("auth_return_url", req.URL.String())
				m.BeginAuthorization(a, r, w, req.Request)
				return nil
			}
			req.SetAuthorizedContext(ctx)
		}
		return h.ServeHTTP(w, req)
	}))
}

// AuthorizeFunc is a convenience function that wraps the given handler closure
// in an authorization middleware.
//
// The Authorize middleware implicitly applies both Authenticate and Sessionize middlewares.
func AuthorizeFunc(m auth.Manager, r model.Role, h func(w http.ResponseWriter, req *route.Request) error) route.Handler {
	return Authorize(m, r, route.HandlerFunc(h))
}
