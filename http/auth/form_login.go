package auth

import (
	"errors"
	"net/http"

	"github.com/motki/core/log"
	"github.com/motki/core/model"
	"github.com/motki/motki-server/http/session"
)

const (
	formLoginCsrfTokenKey = "__motki_form_csrf_token"
	formLoginSessionKey   = "__motki_form_session_id"
)

// CSRFTokenFromSession attempts to retrieve the generated form login CSRF
// token from the user's session storage.
func CSRFTokenFromSession(s *session.Session) (string, bool) {
	return s.String(formLoginCsrfTokenKey)
}

// FormLoginAuthenticator is an authenticator using database-backed users and form logins.
type FormLoginAuthenticator struct {
	loginURL string

	model  *model.Manager
	logger log.Logger
}

// NewFormLoginAuthenticator creates a new authenticator for use with form logins.
func NewFormLoginAuthenticator(m *model.Manager, logger log.Logger, loginURL string) Authenticator {
	logger.Debugf("auth: init form_login authentication")
	logger.Debugf("auth: login URL: %s", loginURL)
	return &FormLoginAuthenticator{
		loginURL: loginURL,
		model:    m,
		logger:   logger,
	}
}

// BeginAuthentication starts the authentication process.
//
// This function generates a CSRF token and redirects the user to the login form.
func (p *FormLoginAuthenticator) BeginAuthentication(s *session.Session, w http.ResponseWriter, req *http.Request) {
	s.NewCSRF(formLoginCsrfTokenKey)
	http.Redirect(w, req, p.loginURL, 302)
}

// FinishAuthentication completes the authentication process.
//
// This function verifies the submitted information against what is stored in the database.
// If the verification passes, an authenticated user session is started.
func (p *FormLoginAuthenticator) FinishAuthentication(s *session.Session, req *http.Request) (*Session, error) {
	if !s.CheckCSRF(formLoginCsrfTokenKey, req.FormValue("_token")) {
		return nil, errors.New("received csrf token did not match expected")
	}
	username := req.FormValue("username")
	password := req.FormValue("password")
	u, key, err := p.model.AuthenticateUser(username, password)
	if err != nil {
		return nil, err
	}
	s.Set(formLoginSessionKey, key)
	s.Set(AuthenticatedUserSessionKey, u.UserID)
	return &Session{
		Session: s,
		user:    u,
	}, nil
}

// VerifyAuthentication checks for an existing authentication in the current session.
//
// This function checks the given session for authentication information and a valid user session.
// If the user session is valid, an authenticated user session is continued.
func (p *FormLoginAuthenticator) VerifyAuthentication(s *session.Session) (*Session, bool) {
	if sid, ok := s.String(formLoginSessionKey); ok {
		u, err := p.model.GetUserBySessionKey(sid)
		if err == nil {
			return &Session{s, u}, true
		}
		p.logger.Warnf("session contains authenticated session key, but couldn't load session: %s", err.Error())
	}
	return nil, false
}
