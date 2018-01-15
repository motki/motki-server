// Package industry contains functionality related to market and industry.
package industry

import (
	"errors"
	"net/http"
	"sort"
	"time"

	"strconv"

	"github.com/motki/core/eveapi"
	"github.com/motki/core/evedb"
	"github.com/motki/core/log"
	"github.com/motki/core/model"
	"github.com/motki/motki-server/http/auth"
	"github.com/motki/motki-server/http/middleware"
	"github.com/motki/motki-server/http/route"
	"github.com/motki/motki-server/http/template"
)

// jobSlice defines how to sort jobs by end date ascending
type jobSlice []*eveapi.IndustryJob

func (s jobSlice) Len() int {
	return len(s)
}

func (s jobSlice) Less(i, j int) bool {
	return s[i].EndDate.After(s[j].EndDate)
}

func (s jobSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type industryModule struct {
	templates template.Renderer
	auth      auth.Manager
	model     *model.Manager
	edb       *evedb.EveDB

	logger log.Logger
}

func New(a auth.Manager, r template.Renderer, mdl *model.Manager, edb *evedb.EveDB, logger log.Logger) *industryModule {
	return &industryModule{
		templates: r,
		auth:      a,
		model:     mdl,
		edb:       edb,

		logger: logger,
	}
}

func (m *industryModule) Init(mux *route.ServeMux) error {
	mux.Handle("/industry/", middleware.AuthorizeFunc(m.auth, model.RoleLogistics, m.indexAction))
	mux.Handle("/industry/blueprints", middleware.AuthorizeFunc(m.auth, model.RoleLogistics, m.blueprintsAction))
	mux.Handle("/industry/structures", middleware.AuthorizeFunc(m.auth, model.RoleLogistics, m.structuresAction))
	mux.Handle("/industry/assets", middleware.AuthorizeFunc(m.auth, model.RoleLogistics, m.assetsAction))
	return nil
}

func (m *industryModule) indexAction(w http.ResponseWriter, req *route.Request) error {
	if req.URL.Path != "/industry/" {
		m.templates.Error(http.StatusNotFound, req, w)
		return nil
	}
	s, ok := req.Auth()
	if !ok {
		m.templates.Error(http.StatusInternalServerError, req, w)
		return errors.New("could not get current authenticated session")
	}
	c, err := m.model.GetCharacter(s.User().CharacterID)
	if err != nil {
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	conf, err := m.model.GetCorporationConfig(c.CorporationID)
	if err != nil {
		m.templates.Error(http.StatusForbidden, req, w)
		return err
	}
	if !conf.OptIn {
		m.templates.Error(http.StatusForbidden, req, w)
		return errors.New("corp is not opted in to data collection")
	}
	a, err := m.model.GetCorporationAuthorization(c.CorporationID)
	if err != nil {
		m.templates.Error(http.StatusForbidden, req, w)
		return err
	}
	jobs, err := m.model.GetCorporationIndustryJobs(a.Context(), c.CorporationID)
	if err != nil {
		if err.Error() == "403 Forbidden" {
			m.templates.Error(http.StatusForbidden, req, w)
			return errors.New("received forbidden response from eve API")
		}
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	sort.Sort(jobSlice(jobs))
	offset := 0
	total := len(jobs)
	const limit = 50
	if page := req.FormValue("page"); page != "" {
		if i, err := strconv.Atoi(page); err == nil {
			if i >= 0 {
				offset = (i - 1) * limit
			}
		}
	}
	if offset > len(jobs) {
		jobs = nil
	} else {
		offplus := offset + limit
		if offplus > len(jobs) {
			offplus = len(jobs)
		}
		jobs = jobs[offset:offplus]
	}
	act := map[int]string{
		1: "Manufacturing",
		3: "Researching TE",
		4: "Researching ME",
		5: "Copying",
		8: "Invention",
		7: "Reverse Engineering",
	}
	m.templates.Render("industry/index.html.twig", req, w, template.Params{"jobs": jobs, "activities": act, "now": time.Now(), "corp_id": c.CorporationID, "page": offset/limit + 1, "total": total})
	return nil
}

func (m *industryModule) blueprintsAction(w http.ResponseWriter, req *route.Request) error {
	s, ok := req.Auth()
	if !ok {
		m.templates.Error(http.StatusInternalServerError, req, w)
		return errors.New("could not get current authenticated session")
	}
	c, err := m.model.GetCharacter(s.User().CharacterID)
	if err != nil {
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	conf, err := m.model.GetCorporationConfig(c.CorporationID)
	if err != nil {
		m.templates.Error(http.StatusForbidden, req, w)
		return err
	}
	if !conf.OptIn {
		m.templates.Error(http.StatusForbidden, req, w)
		return errors.New("corp is not opted in to data collection")
	}
	a, err := m.model.GetCorporationAuthorization(c.CorporationID)
	if err != nil {
		m.templates.Error(http.StatusForbidden, req, w)
		return err
	}
	bps, err := m.model.GetCorporationBlueprints(a.Context(), c.CorporationID)
	if err != nil {
		if err.Error() == "403 Forbidden" {
			m.templates.Error(http.StatusForbidden, req, w)
			return errors.New("received forbidden response from eve API")
		}
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	offset := 0
	total := len(bps)
	const limit = 50
	if page := req.FormValue("page"); page != "" {
		if i, err := strconv.Atoi(page); err == nil {
			if i >= 0 {
				offset = (i - 1) * limit
			}
		}
	}
	if offset > len(bps) {
		bps = nil
	} else {
		offplus := offset + limit
		if offplus > len(bps) {
			offplus = len(bps)
		}
		bps = bps[offset:offplus]
	}
	m.templates.Render("industry/blueprints.html.twig", req, w, template.Params{
		"bps":     bps,
		"corp_id": c.CorporationID,
		"page":    offset/limit + 1,
		"total":   total,
	})
	return nil
}

type helper struct {
	edb   *evedb.EveDB
	model *model.Manager
	cache map[int]string
}

func (h *helper) GetSystem(id int64) string {
	s, err := h.edb.GetSystem(int(id))
	if err != nil {
		return ""
	}
	return s.Name
}

func (h *helper) GetType(id int64) string {
	s, err := h.edb.GetItemType(int(id))
	if err != nil {
		return ""
	}
	return s.Name
}

func (h *helper) GetTypeInt(id int) string {
	s, err := h.edb.GetItemType(id)
	if err != nil {
		return ""
	}
	return s.Name
}

func (h *helper) GetAssetSystem(a *model.Asset) string {
	if v, ok := h.cache[a.LocationID]; ok {
		return v
	}
	s, err := h.model.GetAssetSystem(a)
	if err != nil {
		return ""
	}
	h.cache[a.LocationID] = s.Name
	return s.Name
}

func (m *industryModule) structuresAction(w http.ResponseWriter, req *route.Request) error {
	s, ok := req.Auth()
	if !ok {
		m.templates.Error(http.StatusInternalServerError, req, w)
		return errors.New("could not get current authenticated session")
	}
	c, err := m.model.GetCharacter(s.User().CharacterID)
	if err != nil {
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	conf, err := m.model.GetCorporationConfig(c.CorporationID)
	if err != nil {
		m.templates.Error(http.StatusForbidden, req, w)
		return err
	}
	if !conf.OptIn {
		m.templates.Error(http.StatusForbidden, req, w)
		return errors.New("corp is not opted in to data collection")
	}
	a, err := m.model.GetCorporationAuthorization(c.CorporationID)
	if err != nil {
		m.templates.Error(http.StatusForbidden, req, w)
		return err
	}
	structs, err := m.model.GetCorporationStructures(a.Context(), c.CorporationID)
	if err != nil {
		if err.Error() == "403 Forbidden" {
			m.templates.Error(http.StatusForbidden, req, w)
			return errors.New("received forbidden response from eve API")
		}
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	h := &helper{m.edb, m.model, map[int]string{}}
	offset := 0
	total := len(structs)
	const limit = 50
	if page := req.FormValue("page"); page != "" {
		if i, err := strconv.Atoi(page); err == nil {
			if i >= 0 {
				offset = (i - 1) * limit
			}
		}
	}
	if offset > len(structs) {
		structs = nil
	} else {
		offplus := offset + limit
		if offplus > len(structs) {
			offplus = len(structs)
		}
		structs = structs[offset:offplus]
	}
	m.templates.Render("industry/structures.html.twig", req, w, template.Params{
		"structures": structs,
		"helper":     h,
		"corp_id":    c.CorporationID,
		"page":       offset/limit + 1,
		"total":      total,
	})
	return nil
}

func (m *industryModule) assetsAction(w http.ResponseWriter, req *route.Request) error {
	s, ok := req.Auth()
	if !ok {
		m.templates.Error(http.StatusInternalServerError, req, w)
		return errors.New("could not get current authenticated session")
	}
	c, err := m.model.GetCharacter(s.User().CharacterID)
	if err != nil {
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	conf, err := m.model.GetCorporationConfig(c.CorporationID)
	if err != nil {
		m.templates.Error(http.StatusForbidden, req, w)
		return err
	}
	if !conf.OptIn {
		m.templates.Error(http.StatusForbidden, req, w)
		return errors.New("corp is not opted in to data collection")
	}
	a, err := m.model.GetCorporationAuthorization(c.CorporationID)
	if err != nil {
		m.templates.Error(http.StatusForbidden, req, w)
		return err
	}
	assets, err := m.model.GetCorporationAssets(a.Context(), c.CorporationID)
	if err != nil {
		if err.Error() == "403 Forbidden" {
			m.templates.Error(http.StatusForbidden, req, w)
			return errors.New("received forbidden response from eve API")
		}
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	h := &helper{m.edb, m.model, map[int]string{}}
	offset := 0
	total := len(assets)
	const limit = 50
	if page := req.FormValue("page"); page != "" {
		if i, err := strconv.Atoi(page); err == nil {
			if i >= 0 {
				offset = (i - 1) * limit
			}
		}
	}
	if offset > len(assets) {
		assets = nil
	} else {
		offplus := offset + limit
		if offplus > len(assets) {
			offplus = len(assets)
		}
		assets = assets[offset:offplus]
	}
	m.templates.Render("industry/assets.html.twig", req, w, template.Params{
		"assets":  assets,
		"helper":  h,
		"corp_id": c.CorporationID,
		"page":    offset/limit + 1,
		"total":   total,
	})
	return nil
}
