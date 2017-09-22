// Package auth manages authentication for the motki application.
package auth

import (
	"context"
	"net/http"

	"github.com/motki/motki-server/http/session"
	"github.com/motki/motki/model"
)

const characterIDKey = "__motki_character_id"

// Once a user is authenticated, this key will be populated in the user's
// session. It will contain the current user ID.
const AuthenticatedUserSessionKey = "__motki_user_id"

// An Authenticator implements an authentication mechanism using a user's session.
//
// Generally, an HTTP handler requiring authentication should be wrapped using the
// package level auth.Handle or auth.HandleFunc methods. If the handler does not
// require authorized access to any resources, the role model.RoleAnon can be used
// to skip any further authorization.
type Authenticator interface {
	// BeginAuthentication begins the authentication process.
	BeginAuthentication(s *session.Session, w http.ResponseWriter, r *http.Request)

	// FinishAuthentication completes the authentication process, returning an authenticated session.
	FinishAuthentication(s *session.Session, r *http.Request) (*Session, error)

	// VerifyAuthentication authenticates the session using stored details.
	VerifyAuthentication(s *session.Session) (*Session, bool)
}

// An Authorizer implements an authorization mechanism using a user's session.
//
// Functionality in the Authorizer requires previously authenticated user Sessions.
//
// The general pattern is to use the package level auth.Handle or auth.HandleFunc to
// wrap your HTTP handler in a middleware that will ensure the user is authorized for
// the role specified.
//
// Once authorization is complete, the package level function auth.AuthorizedContext
// can be used inside HTTP handlers to retrieve a Context with necessary token
// information to access a given resource.
type Authorizer interface {
	// BeginAuthorization begins the authorization process for the given role.
	BeginAuthorization(s *Session, r model.Role, w http.ResponseWriter, req *http.Request)

	// FinishAuthorization completes the authorization process, returning an error if it fails.
	FinishAuthorization(s *Session, req *http.Request) error

	// InvalidateAuthorizations removes any associated data for the given role from the session.
	InvalidateAuthorization(s *Session, r model.Role) error

	// AuthorizedContext returns a Context with authorization information inside.
	AuthorizedContext(s *Session, r model.Role) (context.Context, error)
}

// A Manager is a wrapper around an Authenticator, Authorizer, and session manager.
type Manager struct {
	Authenticator
	Authorizer

	Sessions session.Manager
}

// NewManager creates a new Manager, ready for use.
func NewManager(authenticator Authenticator, authorizer Authorizer, sessions session.Manager) Manager {
	return Manager{
		authenticator,
		authorizer,

		sessions,
	}
}

// A Session is an authenticated session.
type Session struct {
	*session.Session

	user *model.User
}

func (s *Session) characterID() int {
	i, ok := s.Int(characterIDKey)
	if !ok {
		return 0
	}
	return i
}

// User returns a normalized, authenticated user for the session.
func (s *Session) User() User {
	return User{s.user, s.characterID()}
}

// A user is a normalized, authenticated user.
type User struct {
	*model.User

	CharacterID int // May be 0 indicating character not yet selected.
}
