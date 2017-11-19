// Package home contains the basic public functionality of the MOTKI app.
package home

import (
	"net/http"

	"github.com/motki/motki-server/http/middleware"
	"github.com/motki/motki-server/http/route"
	"github.com/motki/motki-server/http/session"
	"github.com/motki/motki-server/http/template"
	"github.com/motki/motki-server/mail"
	"github.com/motki/motki/log"
)

const (
	sessionKeyUnsubCSRFToken    = "__motki_unsubscribe_csrf"
	sessionKeyUnsubErrorMessage = "__motki_unsubscribe_error"
)

type homeModule struct {
	sessions  session.Manager
	templates template.Renderer
	mailer    *mail.Sender

	logger log.Logger
}

func New(s session.Manager, r template.Renderer, m *mail.Sender, l log.Logger) *homeModule {
	return &homeModule{
		sessions:  s,
		templates: r,
		mailer:    m,

		logger: l,
	}
}

func (m *homeModule) Init(mux *route.ServeMux) error {
	mux.Handle("/", middleware.SessionizeFunc(m.sessions, m.indexAction))
	mux.Handle("/about", middleware.SessionizeFunc(m.sessions, m.aboutAction))
	mux.Handle("/press", middleware.SessionizeFunc(m.sessions, m.pressAction))
	mux.Handle("/press/yc119/h1", middleware.SessionizeFunc(m.sessions, m.press2Action))
	mux.Handle("/recruitment", middleware.SessionizeFunc(m.sessions, m.recruitmentAction))
	mux.Handle("/unsubscribe", middleware.SessionizeFunc(m.sessions, m.unsubscribeAction))
	mux.Handle("/privacy", middleware.SessionizeFunc(m.sessions, m.privacyAction))
	return nil
}

func (m *homeModule) indexAction(w http.ResponseWriter, req *route.Request) error {
	if req.URL.Path != "/" {
		m.templates.Error(http.StatusNotFound, req, w)
		return nil
	}
	m.templates.Render("home/index.html.twig", req, w, nil)
	return nil
}

func (m *homeModule) aboutAction(w http.ResponseWriter, req *route.Request) error {
	m.templates.Render("home/about.html.twig", req, w, nil)
	return nil
}

func (m *homeModule) pressAction(w http.ResponseWriter, req *route.Request) error {
	m.templates.Render("home/press.html.twig", req, w, nil)
	return nil
}

func (m *homeModule) press2Action(w http.ResponseWriter, req *route.Request) error {
	m.templates.Render("home/press2.html.twig", req, w, nil)
	return nil
}

func (m *homeModule) recruitmentAction(w http.ResponseWriter, req *route.Request) error {
	m.templates.Render("home/recruitment.html.twig", req, w, nil)
	return nil
}

func (m *homeModule) unsubscribeAction(w http.ResponseWriter, req *route.Request) error {
	sess, ok := req.Session()
	if !ok {
		m.logger.Warnf("woops, could not get current session from context")
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	// TODO: This could be seriously refactored
	if req.Method == http.MethodPost {
		if sess.CheckCSRF(sessionKeyUnsubCSRFToken, req.FormValue("_token")) {
			email := req.FormValue("email")
			if email != "" {
				err := m.mailer.DoNotSend.Add(mail.Recipient{Email: email})
				if err != nil {
					m.logger.Warnf("woops, could not add someone to the unsubscribe: %s", err.Error())
					m.templates.Error(http.StatusInternalServerError, req, w)
					return err
				}
				m.templates.Render("home/unsubscribed.html.twig", req, w, nil)
				return nil
			}
			sess.SetFlash(sessionKeyUnsubErrorMessage, "email cannot be blank")
		} else {
			sess.SetFlash(sessionKeyUnsubErrorMessage, "invalid csrf token")
		}
		http.Redirect(w, req.Request, "/unsubscribe", http.StatusFound)
		return nil
	}
	m.templates.Render("home/unsubscribe.html.twig", req, w, template.Params{
		"error":      sess.Flash(sessionKeyUnsubErrorMessage),
		"email":      req.FormValue("email"),
		"csrf_token": sess.NewCSRF(sessionKeyUnsubCSRFToken),
	})
	return nil
}

func (m *homeModule) privacyAction(w http.ResponseWriter, req *route.Request) error {
	m.templates.Render("home/privacy.html.twig", req, w, nil)
	return nil
}
