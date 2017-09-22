package account

import (
	"net/http"

	"github.com/motki/motki/evedb"
	"github.com/motki/motki/log"
	"github.com/motki/motki/model"
	"github.com/motki/motki-server/http/auth"
	"github.com/motki/motki-server/http/middleware"
	"github.com/motki/motki-server/http/route"
	"github.com/motki/motki-server/http/template"
)

const (
	sessionKeyUnlinkMainCSRF = "__motki_account_unlink_main_csrf"
	sessionKeyUnlinkLogiCSRF = "__motki_account_unlink_logi_csrf"
)

type Character struct {
	model.Character

	Corp     model.Corporation
	Alliance model.Alliance

	Race      evedb.Race
	Bloodline evedb.Bloodline
	Ancestry  evedb.Ancestry
}

type accountModule struct {
	auth      auth.Manager
	templates template.Renderer
	model     *model.Manager
	evedb     *evedb.EveDB

	logger log.Logger
}

func New(a auth.Manager, r template.Renderer, m *model.Manager, edb *evedb.EveDB, logger log.Logger) *accountModule {
	return &accountModule{
		auth:      a,
		templates: r,
		model:     m,
		evedb:     edb,

		logger: logger,
	}
}

func (m *accountModule) getCharacter(characterID int) (char Character, err error) {
	c, err := m.model.GetCharacter(characterID)
	if err != nil {
		return char, err
	}
	char.Character = *c
	corp, err := m.model.GetCorporation(c.CorporationID)
	if err != nil {
		return char, err
	}
	char.Corp = *corp
	var alliance *model.Alliance
	if corp.AllianceID != 0 {
		alliance, err = m.model.GetAlliance(corp.AllianceID)
		if err != nil {
			return char, err
		}
		char.Alliance = *alliance
	}
	race, err := m.evedb.GetRace(c.RaceID)
	if err != nil {
		return char, err
	}
	char.Race = *race
	ancestry, err := m.evedb.GetAncestry(c.AncestryID)
	if err != nil {
		return char, err
	}
	char.Ancestry = *ancestry
	bloodline, err := m.evedb.GetBloodline(c.BloodlineID)
	char.Bloodline = *bloodline
	if err != nil {
		return char, err
	}
	return char, nil
}

func (m *accountModule) Init(mux *route.ServeMux) error {
	mux.Handle("/account/", middleware.AuthorizeFunc(m.auth, model.RoleUser, m.indexAction))
	mux.Handle("/account/my-characters", middleware.AuthorizeFunc(m.auth, model.RoleUser, m.charactersAction))

	mux.Handle("/account/unlink-main", middleware.AuthenticateFunc(m.auth, m.unlinkMainAction))
	mux.Handle("/account/link-main", middleware.AuthorizeFunc(m.auth, model.RoleUser, m.linkMainAction))

	mux.Handle("/account/unlink-logistics", middleware.AuthenticateFunc(m.auth, m.unlinkLogisticsAction))
	mux.Handle("/account/link-logistics", middleware.AuthorizeFunc(m.auth, model.RoleLogistics, m.linkLogisticsAction))
	return nil
}

func (m *accountModule) indexAction(w http.ResponseWriter, r *route.Request) error {
	if r.URL.Path != "/account/" {
		m.templates.Error(http.StatusNotFound, r, w)
		return nil
	}
	s, ok := r.Auth()
	if !ok {
		m.logger.Warnf("woops, could not get authenticated session from context")
		m.templates.Error(http.StatusInternalServerError, r, w)
		return nil
	}
	char, err := m.getCharacter(s.User().CharacterID)
	if err != nil {
		m.logger.Warnf("woops, could not get char info: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}
	m.templates.Render("account/index.html.twig", r, w, template.Params{
		"character": char,
	})
	return nil
}

func (m *accountModule) charactersAction(w http.ResponseWriter, r *route.Request) error {
	s, ok := r.Auth()
	if !ok {
		m.logger.Warnf("woops, could not get authenticated session from context")
		m.templates.Error(http.StatusInternalServerError, r, w)
		return nil
	}
	char, err := m.getCharacter(s.User().CharacterID)
	if err != nil {
		m.logger.Warnf("woops, could not get char info: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}
	var logi Character
	if logiAuth, err := m.model.GetAuthorization(s.User().User, model.RoleLogistics); err == nil {
		logi, err = m.getCharacter(logiAuth.CharacterID)
		if err != nil {
			m.logger.Warnf("woops, could not get char info: %s", err.Error())
			m.templates.Error(http.StatusInternalServerError, r, w)
			return err
		}
	} else {
		m.logger.Debugf("woops, failed to get user's logi char: %s", err.Error())
	}
	mainCsrf := s.NewCSRF(sessionKeyUnlinkMainCSRF)
	logiCsrf := s.NewCSRF(sessionKeyUnlinkLogiCSRF)
	m.templates.Render("account/characters.html.twig", r, w, template.Params{
		"main":      char,
		"main_csrf": mainCsrf,
		"logistics": logi,
		"logi_csrf": logiCsrf,
	})
	return nil
}

func (m *accountModule) unlinkMainAction(w http.ResponseWriter, r *route.Request) error {
	if r.Method != http.MethodPost {
		m.logger.Debugf("woops, did not receive expected POST method request")
		m.templates.Error(http.StatusMethodNotAllowed, r, w)
		return nil
	}
	s, ok := r.Auth()
	if !ok {
		m.logger.Warnf("woops, could not get authenticated session from context")
		m.templates.Error(http.StatusInternalServerError, r, w)
		return nil
	}
	if !s.CheckCSRF(sessionKeyUnlinkMainCSRF, r.FormValue("_token")) {
		m.logger.Warnf("woops, csrf did not match")
		m.templates.Error(http.StatusInternalServerError, r, w)
		return nil
	}
	err := m.model.RemoveAuthorization(s.User().User, model.RoleUser)
	if err != nil {
		m.logger.Warnf("woops, could not remove authorization: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}
	err = m.auth.InvalidateAuthorization(s, model.RoleUser)
	if err != nil {
		m.logger.Warnf("woops, could not invalidate authorization: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}
	http.Redirect(w, r.Request, "/account/link-main", http.StatusFound)
	return nil
}

func (m *accountModule) linkMainAction(w http.ResponseWriter, r *route.Request) error {
	http.Redirect(w, r.Request, "/account/my-characters", http.StatusFound)
	return nil
}

func (m *accountModule) unlinkLogisticsAction(w http.ResponseWriter, r *route.Request) error {
	if r.Method != http.MethodPost {
		m.logger.Debugf("woops, did not receive expected POST method request")
		m.templates.Error(http.StatusMethodNotAllowed, r, w)
		return nil
	}
	s, ok := r.Auth()
	if !ok {
		m.logger.Warnf("woops, could not get authenticated session from context")
		m.templates.Error(http.StatusInternalServerError, r, w)
		return nil
	}
	if !s.CheckCSRF(sessionKeyUnlinkLogiCSRF, r.FormValue("_token")) {
		m.logger.Warnf("woops, csrf did not match")
		m.templates.Error(http.StatusInternalServerError, r, w)
		return nil
	}
	err := m.model.RemoveAuthorization(s.User().User, model.RoleLogistics)
	if err != nil {
		m.logger.Warnf("woops, could not remove authorization: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}
	err = m.auth.InvalidateAuthorization(s, model.RoleLogistics)
	if err != nil {
		m.logger.Warnf("woops, could not invalidate authorization: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}
	http.Redirect(w, r.Request, "/account/link-logistics", http.StatusFound)
	return nil
}

func (m *accountModule) linkLogisticsAction(w http.ResponseWriter, r *route.Request) error {
	http.Redirect(w, r.Request, "/account/my-characters", http.StatusFound)
	return nil
}
