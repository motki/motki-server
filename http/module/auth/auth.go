package auth

import (
	"bytes"
	"encoding/base64"
	"net/http"

	"github.com/motki/motki/log"
	"github.com/motki/motki/model"
	"github.com/motki/motki/worker"
	"github.com/motki/motki-server/http/auth"
	"github.com/motki/motki-server/http/middleware"
	"github.com/motki/motki-server/http/route"
	"github.com/motki/motki-server/http/session"
	"github.com/motki/motki-server/http/template"
	"github.com/motki/motki-server/mail"
)

const (
	sessKeyRegisterCSRFToken    = "__motki_register_csrf"
	sessKeyLoginErrorMessage    = "__motki_login_error"
	sessKeyRegisterErrorMessage = "__motki_register_error"
)

type authModule struct {
	sessions  session.Manager
	templates template.Renderer
	auth      auth.Manager
	model     *model.Manager
	work      *worker.Scheduler
	mailer    *mail.Sender

	logger log.Logger
}

func New(s session.Manager, a auth.Manager, r template.Renderer, m *model.Manager, w *worker.Scheduler, e *mail.Sender, logger log.Logger) *authModule {
	return &authModule{
		sessions:  s,
		auth:      a,
		templates: r,
		model:     m,
		work:      w,
		mailer:    e,

		logger: logger,
	}
}

func (m *authModule) Init(mux *route.ServeMux) error {
	mux.Handle("/login", middleware.AuthenticateFunc(m.auth, m.beginLoginHandler))
	mux.Handle("/login/begin", middleware.SessionizeFunc(m.sessions, m.formLoginHandler))
	mux.Handle("/login/finish", middleware.SessionizeFunc(m.sessions, m.finishLoginHandler))

	mux.Handle("/auth/finish-login", middleware.AuthenticateFunc(m.auth, m.finishAuthorizationHandler))

	mux.Handle("/register", middleware.SessionizeFunc(m.sessions, m.registerHandler))
	mux.Handle("/verify", middleware.SessionizeFunc(m.sessions, m.verifyHandler))

	mux.Handle("/logout", middleware.SessionizeFunc(m.sessions, m.logoutHandler))
	return nil
}

func (m *authModule) formLoginHandler(w http.ResponseWriter, req *route.Request) error {
	s, ok := req.Session()
	if !ok {
		m.logger.Warnf("woops, could not get current session from context")
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	if _, ok := m.auth.VerifyAuthentication(s); ok {
		// Already logged in.
		http.Redirect(w, req.Request, "/account/", http.StatusFound)
		return nil
	}
	tok, ok := auth.CSRFTokenFromSession(s)
	if !ok {
		m.logger.Warnf("woops, could not get CSRF token from context")
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	regTok := s.NewCSRF(sessKeyRegisterCSRFToken)
	m.templates.Render("auth/login.html.twig", req, w, template.Params{
		"login_error":    s.Flash(sessKeyLoginErrorMessage),
		"register_error": s.Flash(sessKeyRegisterErrorMessage),
		"csrf_token":     tok,
		"register_token": regTok,
	})
	return nil
}

func (m *authModule) beginLoginHandler(w http.ResponseWriter, req *route.Request) error {
	_, ok := req.Auth()
	if !ok {
		m.logger.Warnf("woops, could not get current authenticated session from context")
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	// Already logged in.
	http.Redirect(w, req.Request, "/account/", http.StatusFound)
	return nil
}

func (m *authModule) finishLoginHandler(w http.ResponseWriter, req *route.Request) error {
	if req.Method != http.MethodPost {
		http.Redirect(w, req.Request, "/login", http.StatusFound)
		return nil
	}
	s, ok := req.Session()
	if !ok {
		m.logger.Warnf("woops, could not get current session from context")
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	_, err := m.auth.FinishAuthentication(s, req.Request)
	if err != nil {
		m.logger.Debugf("error authenticating: %s", err.Error())
		s.SetFlash(sessKeyLoginErrorMessage, "invalid username or password")
		http.Redirect(w, req.Request, "/login", http.StatusFound)
		return err
	}
	http.Redirect(w, req.Request, "/account/", http.StatusFound)
	return nil
}

func (m *authModule) finishAuthorizationHandler(w http.ResponseWriter, req *route.Request) error {
	sess, ok := req.Auth()
	if !ok {
		m.logger.Warnf("woops, could not get current authenticated session from context")
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	err := m.auth.FinishAuthorization(sess, req.Request)
	if err != nil {
		m.logger.Warnf("woops, failed to finish authorization: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	if url, ok := sess.String("auth_return_url"); ok {
		sess.Remove("auth_return_url")
		http.Redirect(w, req.Request, url, http.StatusFound)
		return nil
	}
	http.Redirect(w, req.Request, "/account/", http.StatusFound)
	return nil
}

func (m *authModule) registerHandler(w http.ResponseWriter, req *route.Request) error {
	if req.Method != http.MethodPost {
		m.logger.Debugf("woops, did not receive expected POST method request")
		m.templates.Error(http.StatusMethodNotAllowed, req, w)
		return nil
	}
	sess, ok := req.Session()
	if !ok {
		m.logger.Warnf("woops, could not get current session from context")
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	if !sess.CheckCSRF(sessKeyRegisterCSRFToken, req.FormValue("_token")) {
		m.logger.Debugf("csrf token mismatch")
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	username := req.FormValue("username")
	email := req.FormValue("email")
	password := req.FormValue("password")
	user, err := m.model.NewUser(username, email, password)
	if err != nil {
		if err == model.ErrUserExists {
			sess.SetFlash(sessKeyRegisterErrorMessage, "username or email is taken")
			http.Redirect(w, req.Request, "/login", http.StatusFound)
			return nil
		} else if err == model.ErrMissingField {
			sess.SetFlash(sessKeyRegisterErrorMessage, "all fields are required")
			http.Redirect(w, req.Request, "/login", http.StatusFound)
			return nil
		}
		m.logger.Warnf("woops, error creating user: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	m.work.ScheduleFunc(func() error {
		h, err := m.model.CreateUserVerificationHash(user)
		if err != nil {
			return err
		}
		rec := mail.Recipient{
			Name:  username,
			Email: email,
		}
		b := &bytes.Buffer{}
		err = m.templates.Render("mail/verify_email.html.twig", nil, b, template.Params{
			"username": username,
			"email":    email,
			"hash":     base64.RawURLEncoding.EncodeToString(h),
		})
		if err != nil {
			return err
		}
		return m.mailer.Send(rec, "Verify your email address", b.String())
	})
	m.templates.Render("auth/registered.html.twig", req, w, template.Params{
		"user": user,
	})
	return nil
}

func (m *authModule) verifyHandler(w http.ResponseWriter, req *route.Request) error {
	email := req.FormValue("email")
	hash := req.FormValue("hash")
	if email == "" || hash == "" {
		m.logger.Debugf("email or hash is empty, unable to verify user")
		m.templates.Error(http.StatusBadRequest, req, w)
		return nil
	}
	h, err := base64.RawURLEncoding.DecodeString(hash)
	if err != nil {
		m.logger.Warnf("unable to decode base64 encoded hash: %s", err.Error())
		m.templates.Error(http.StatusBadRequest, req, w)
		return err
	}
	ok, err := m.model.VerifyUserEmail(email, h)
	if !ok {
		m.logger.Debugf("unable to verify user, possible error: %v", err)
		m.templates.Error(http.StatusBadRequest, req, w)
		return err
	}
	m.templates.Render("auth/verified.html.twig", req, w, nil)
	return nil
}

func (m *authModule) logoutHandler(w http.ResponseWriter, req *route.Request) error {
	m.sessions.Invalidate(req.Request, w)
	http.Redirect(w, req.Request, "/", 302)
	return nil
}
