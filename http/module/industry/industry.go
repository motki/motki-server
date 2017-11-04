// Package industry contains functionality related to market and industry.
package industry

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/motki/motki-server/http/auth"
	"github.com/motki/motki-server/http/middleware"
	"github.com/motki/motki-server/http/route"
	"github.com/motki/motki-server/http/template"
	"github.com/motki/motki/eveapi"
	"github.com/motki/motki/evedb"
	"github.com/motki/motki/log"
	"github.com/motki/motki/model"
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
		m.logger.Warnf("woops, could not get current authenticated session from  context")
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	c, err := m.model.GetCharacter(s.User().CharacterID)
	if err != nil {
		m.logger.Warnf("woops, could not get char info: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	apiCtx, ok := req.AuthorizedContext()
	if !ok {
		m.logger.Warnf("woops, could not get authorized context")
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	jobs, err := m.model.GetCorporationIndustryJobs(apiCtx, c.CorporationID)
	if err != nil {
		if err.Error() == "403 Forbidden" {
			m.templates.Error(http.StatusForbidden, req, w)
			return nil
		}
		m.logger.Warnf("woops, failed to get corp jobs: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	sort.Sort(jobSlice(jobs))
	act := map[int]string{
		1: "Manufacturing",
		3: "Researching TE",
		4: "Researching ME",
		5: "Copying",
		8: "Invention",
		7: "Reverse Engineering",
	}
	m.templates.Render("industry/index.html.twig", req, w, template.Params{"jobs": jobs, "activities": act, "now": time.Now()})
	return nil
}

func (m *industryModule) blueprintsAction(w http.ResponseWriter, req *route.Request) error {
	s, ok := req.Auth()
	if !ok {
		m.logger.Warnf("woops, could not get current authenticated session")
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	c, err := m.model.GetCharacter(s.User().CharacterID)
	if err != nil {
		m.logger.Warnf("woops, failed to get char info: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	apiCtx, ok := req.AuthorizedContext()
	if !ok {
		m.logger.Warnf("woops, failed to get api context: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	bps, err := m.model.GetCorporationBlueprints(apiCtx, c.CorporationID)
	if err != nil {
		if err.Error() == "403 Forbidden" {
			m.templates.Error(http.StatusForbidden, req, w)
			return nil
		}
		m.logger.Warnf("woops, failed to get corp blueprints: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	m.templates.Render("industry/blueprints.html.twig", req, w, template.Params{"bps": bps})
	return nil
}

type helper struct {
	edb   *evedb.EveDB
	model *model.Manager
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
	s, err := h.model.GetAssetSystem(a)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return s.Name
}

func (m *industryModule) structuresAction(w http.ResponseWriter, req *route.Request) error {
	s, ok := req.Auth()
	if !ok {
		m.logger.Warnf("woops, could not get current authenticated session")
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	c, err := m.model.GetCharacter(s.User().CharacterID)
	if err != nil {
		m.logger.Warnf("woops, failed to get char info: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	apiCtx, ok := req.AuthorizedContext()
	if !ok {
		m.logger.Warnf("woops, failed to get corp jobs: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	structs, err := m.model.GetCorporationStructures(apiCtx, c.CorporationID)
	if err != nil {
		if err.Error() == "403 Forbidden" {
			m.templates.Error(http.StatusForbidden, req, w)
			return nil
		}
		m.logger.Warnf("woops, failed to get structures: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	h := &helper{m.edb, m.model}
	m.templates.Render("industry/structures.html.twig", req, w, template.Params{
		"structures": structs,
		"helper":     h,
	})
	return nil
}

func (m *industryModule) assetsAction(w http.ResponseWriter, req *route.Request) error {
	s, ok := req.Auth()
	if !ok {
		m.logger.Warnf("woops, could not get current authenticated session")
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	c, err := m.model.GetCharacter(s.User().CharacterID)
	if err != nil {
		m.logger.Warnf("woops, failed to get char info: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	apiCtx, ok := req.AuthorizedContext()
	if !ok {
		m.logger.Warnf("woops, failed to get corp jobs: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, req, w)
		return nil
	}
	assets, err := m.model.GetCorporationAssets(apiCtx, c.CorporationID)
	if err != nil {
		if err.Error() == "403 Forbidden" {
			m.templates.Error(http.StatusForbidden, req, w)
			return nil
		}
		m.logger.Warnf("woops, failed to get structures: %s", err.Error())
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	spew.Dump(m.model.FetchCorporationDetail(apiCtx))
	spew.Dump(m.model.GetCorporationOrders(apiCtx, c.CorporationID))
	spew.Dump(m.model.GetCorporationOrder(apiCtx, c.CorporationID, 4976116595))
	return nil
	h := &helper{m.edb, m.model}
	m.templates.Render("industry/assets.html.twig", req, w, template.Params{
		"assets": assets,
		"helper": h,
	})
	return nil
}
