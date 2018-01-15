// Package market contains functionality related to market and industry.
package market

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/motki/core/evedb"
	"github.com/motki/core/log"
	"github.com/motki/core/proto/client"
	"github.com/motki/motki-server/http/auth"
	"github.com/motki/motki-server/http/middleware"
	"github.com/motki/motki-server/http/route"
	"github.com/motki/motki-server/http/template"
)

const (
	sessionKeyIndexCSRF     = "__motki_market_csrf"
	sessionKeyIndexTypeIDs  = "__motki_market_type_ids"
	sessionKeyIndexRuns     = "__motki_market_runs"
	sessionKeyIndexRegionID = "__motki_market_region_id"

	sessionKeyLookupCSRF = "__motki_market_lookup_csrf"
)

type marketModule struct {
	auth      auth.Manager
	templates template.Renderer
	client    client.Client

	logger log.Logger
}

func New(a auth.Manager, r template.Renderer, cl client.Client, logger log.Logger) *marketModule {
	return &marketModule{
		auth:      a,
		templates: r,
		client:    cl,
		logger:    logger,
	}
}

func (m *marketModule) Init(mux *route.ServeMux) error {
	mux.Handle("/market", middleware.AuthenticateFunc(m.auth, m.indexAction))
	mux.Handle("/market/lookup", middleware.AuthenticateFunc(m.auth, m.lookupAction))
	return nil
}

func (m *marketModule) indexAction(w http.ResponseWriter, req *route.Request) error {
	sess, ok := req.Session()
	if !ok {
		m.templates.Error(http.StatusBadRequest, req, w)
		return errors.New("expected request to have a session")
	}
	ids := make([]int, 0)
	runs := 1
	regionID := 0
	if req.Method == http.MethodPost {
		if !sess.CheckCSRF(sessionKeyIndexCSRF, req.FormValue("_token")) {
			m.templates.Error(http.StatusBadRequest, req, w)
			return errors.New("csrf mismatch")
		}
		if err := req.ParseForm(); err != nil {
			m.logger.Debugf("unable to parse form: %s", err.Error())
		}
		typeIDs := req.Form["type_id"]
		for _, id := range typeIDs {
			if i, err := strconv.Atoi(id); err == nil {
				ids = append(ids, i)
			}
		}
		rp := req.Form["runs"]
		if len(rp) > 0 {
			if r, err := strconv.Atoi(rp[0]); err == nil {
				runs = r
			}
		}
		rg := req.Form["region"]
		if len(rg) > 0 {
			if r, err := strconv.Atoi(rg[0]); err == nil {
				regionID = r
			}
		}
		sess.Set(sessionKeyIndexTypeIDs, ids)
		sess.Set(sessionKeyIndexRuns, runs)
		sess.Set(sessionKeyIndexRegionID, regionID)
		http.Redirect(w, req.Request, req.URL.String(), http.StatusFound)
		return nil
	} else {
		if v, ok := sess.Get(sessionKeyIndexTypeIDs); ok {
			if vs, ok := v.([]interface{}); ok {
				for _, vsi := range vs {
					if id, ok := vsi.(float64); ok {
						ids = append(ids, int(id))
					}
				}
			}
		}
		if v, ok := sess.Get(sessionKeyIndexRuns); ok {
			if vr, ok := v.(float64); ok {
				runs = int(vr)
			}
		}
		if v, ok := sess.Get(sessionKeyIndexRegionID); ok {
			if vr, ok := v.(float64); ok {
				regionID = int(vr)
			}
		}
	}
	if len(ids) == 0 {
		ids = append(ids, 16242, 4393, 25266, 32880)
	}
	if regionID == 0 {
		regionID = 10000002 // The Forge
	}
	var bps []*evedb.MaterialSheet
	for _, id := range ids {
		bp, err := m.client.GetMaterialSheet(id)
		if err != nil {
			m.templates.Error(http.StatusInternalServerError, req, w)
			return err
		}
		bps = append(bps, bp)
	}
	mats := map[int]struct{}{}
	for _, bp := range bps {
		for _, mat := range bp.Materials {
			mats[mat.ID] = struct{}{}
		}
	}
	matIDs := []int{}
	for matID := range mats {
		matIDs = append(matIDs, matID)
	}
	ids = append(ids, matIDs...)
	stats, err := m.client.GetMarketPrices(ids[0], ids[1:]...)
	if err != nil {
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	avgPrices := map[int]float64{}
	for _, s := range stats {
		avgPrices[s.TypeID], _ = s.Avg.Float64()
	}
	regions, err := m.client.GetRegions()
	if err != nil {
		m.templates.Error(http.StatusInternalServerError, req, w)
		return err
	}
	m.templates.Render("market/index.html.twig", req, w, template.Params{
		"blueprints":  bps,
		"prices":      avgPrices,
		"runs":        runs,
		"regionID":    regionID,
		"regions":     regions,
		"csrf_token":  sess.NewCSRF(sessionKeyIndexCSRF),
		"lookup_csrf": sess.NewCSRF(sessionKeyLookupCSRF),
	})
	return nil
}

func (m *marketModule) lookupAction(w http.ResponseWriter, req *route.Request) error {
	sess, ok := req.Session()
	if !ok {
		renderJSON(http.StatusBadRequest, w, nil)
		return errors.New("expected request to have session")
	}
	payload := struct {
		Items []*evedb.ItemType `json:"items"`
		CSRF  string            `json:"csrf"`
	}{
		make([]*evedb.ItemType, 0),
		"",
	}
	tk := req.URL.Query()["_token"]
	if len(tk) == 0 || !sess.CheckCSRF(sessionKeyLookupCSRF, tk[0]) {
		payload.CSRF = sess.NewCSRF(sessionKeyLookupCSRF)
		renderJSON(http.StatusOK, w, payload)
		return errors.New("csrf mismatch")
	}
	payload.CSRF = sess.NewCSRF(sessionKeyLookupCSRF)
	var query string
	qp := req.URL.Query()["query"]
	if len(qp) == 0 {
		renderJSON(http.StatusOK, w, payload)
		return nil
	} else {
		query = qp[0]
	}
	items, err := m.client.QueryItemTypes(query, evedb.InterestingItemCategories...)
	if err == nil {
		payload.Items = items
	}
	renderJSON(http.StatusOK, w, payload)
	return err
}

func renderJSON(code int, w http.ResponseWriter, payload interface{}) {
	w.WriteHeader(code)
	b, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}
