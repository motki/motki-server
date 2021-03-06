package account

import (
	"net/http"

	"time"

	"github.com/motki/core/evedb"
	"github.com/motki/core/log"
	"github.com/motki/core/model"
	"github.com/motki/core/worker"
	"github.com/motki/motki-server/http/auth"
	"github.com/motki/motki-server/http/middleware"
	"github.com/motki/motki-server/http/route"
	"github.com/motki/motki-server/http/template"
)

const (
	sessionKeyUnlinkMainCSRF     = "__motki_account_unlink_main_csrf"
	sessionKeyUnlinkLogiCSRF     = "__motki_account_unlink_logi_csrf"
	sessionKeyUnlinkDirectorCSRF = "__motki_account_unlink_director_csrf"
	sessionKeyUpdateCorpCSRF     = "__motki_account_update_corp_csrf"
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
	work      *worker.Scheduler

	logger log.Logger
}

func New(a auth.Manager, r template.Renderer, m *model.Manager, edb *evedb.EveDB, work *worker.Scheduler, logger log.Logger) *accountModule {
	return &accountModule{
		auth:      a,
		templates: r,
		model:     m,
		evedb:     edb,
		work:      work,

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

	mux.Handle("/account/unlink-director", middleware.AuthenticateFunc(m.auth, m.unlinkDirectorAction))
	mux.Handle("/account/link-director", middleware.AuthorizeFunc(m.auth, model.RoleDirector, m.linkDirectorAction))

	mux.Handle("/account/manage-corp", middleware.AuthorizeFunc(m.auth, model.RoleDirector, m.editCorpAction))
	mux.Handle("/account/manage-corp/update", middleware.AuthorizeFunc(m.auth, model.RoleDirector, m.updateCorpAction))
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
	var director Character
	if directorAuth, err := m.model.GetAuthorization(s.User().User, model.RoleDirector); err == nil {
		director, err = m.getCharacter(directorAuth.CharacterID)
		if err != nil {
			m.logger.Warnf("woops, could not get char info: %s", err.Error())
			m.templates.Error(http.StatusInternalServerError, r, w)
			return err
		}
	} else {
		m.logger.Debugf("woops, failed to get user's director char: %s", err.Error())
	}
	mainCsrf := s.NewCSRF(sessionKeyUnlinkMainCSRF)
	logiCsrf := s.NewCSRF(sessionKeyUnlinkLogiCSRF)
	directorCsrf := s.NewCSRF(sessionKeyUnlinkDirectorCSRF)
	m.templates.Render("account/characters.html.twig", r, w, template.Params{
		"main":          char,
		"main_csrf":     mainCsrf,
		"logistics":     logi,
		"logi_csrf":     logiCsrf,
		"director":      director,
		"director_csrf": directorCsrf,
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

func (m *accountModule) unlinkDirectorAction(w http.ResponseWriter, r *route.Request) error {
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
	err := m.model.RemoveAuthorization(s.User().User, model.RoleDirector)
	if err != nil {
		m.logger.Warnf("woops, could not remove authorization: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}
	err = m.auth.InvalidateAuthorization(s, model.RoleDirector)
	if err != nil {
		m.logger.Warnf("woops, could not invalidate authorization: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}
	http.Redirect(w, r.Request, "/account/link-director", http.StatusFound)
	return nil
}

func (m *accountModule) linkDirectorAction(w http.ResponseWriter, r *route.Request) error {
	http.Redirect(w, r.Request, "/account/my-characters", http.StatusFound)
	return nil
}

func (m *accountModule) editCorpAction(w http.ResponseWriter, r *route.Request) error {
	s, ok := r.Auth()
	if !ok {
		m.logger.Warnf("woops, could not get authenticated session from context")
		m.templates.Error(http.StatusInternalServerError, r, w)
		return nil
	}
	auth, err := m.model.GetAuthorization(s.User().User, model.RoleDirector)
	if err != nil {
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}
	corp, err := m.model.GetCorporation(auth.CorporationID)
	if err != nil {
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}
	detail, err := m.model.GetCorporationDetail(auth.CorporationID)
	if err != nil && err != model.ErrCorpNotRegistered {
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}
	config, err := m.model.GetCorporationConfig(auth.CorporationID)
	if err != nil {
		if err != model.ErrCorpNotRegistered {
			m.templates.Error(http.StatusInternalServerError, r, w)
			return err
		}
		config = &model.CorporationConfig{}
	}
	csrfToken := s.NewCSRF(sessionKeyUpdateCorpCSRF)
	m.templates.Render("account/edit_corp.html.twig", r, w, template.Params{
		"corp":       corp,
		"config":     config,
		"detail":     detail,
		"csrf_token": csrfToken,
	})
	return nil
}

func (m *accountModule) updateCorpAction(w http.ResponseWriter, r *route.Request) error {
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
	if !s.CheckCSRF(sessionKeyUpdateCorpCSRF, r.FormValue("_token")) {
		m.logger.Warnf("woops, csrf did not match")
		m.templates.Error(http.StatusInternalServerError, r, w)
		return nil
	}
	a, err := m.model.GetAuthorization(s.User().User, model.RoleDirector)
	if err != nil {
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}
	config, err := m.model.GetCorporationConfig(a.CorporationID)
	if err != nil {
		if err != model.ErrCorpNotRegistered {
			m.templates.Error(http.StatusInternalServerError, r, w)
			return err
		}
		config = &model.CorporationConfig{
			CreatedBy: s.User().UserID,
			CreatedAt: time.Now(),
		}
	}
	if r.PostForm.Get("opt_in") == "1" {
		if !config.OptIn {
			config.OptIn = true
			config.OptInDate = time.Now()
			config.OptInBy = s.User().UserID
		}
	} else if config.OptIn {
		config.OptIn = false
		config.OptInDate = time.Now()
		config.OptInBy = s.User().UserID
	}
	err = m.model.SaveCorporationConfig(a.CorporationID, config)
	if err != nil {
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}
	_, err = m.model.FetchCorporationDetail(a.Context())
	if err != nil {
		m.templates.Error(http.StatusInternalServerError, r, w)
		return err
	}

	http.Redirect(w, r.Request, "/account/manage-corp", http.StatusFound)
	return nil
}
