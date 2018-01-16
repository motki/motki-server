package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"

	"github.com/motki/core/eveapi"
	"github.com/motki/core/log"
)

var ErrCorpNotRegistered = errors.New("ceo or director is not registered for the given corporation")

type Corporation struct {
	CorporationID int
	Name          string
	AllianceID    int
	CreationDate  time.Time
	Description   string
	Ticker        string
}

func (m *Manager) GetCorporation(corporationID int) (*Corporation, error) {
	c, err := m.getCorporationFromDB(corporationID)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return m.getCorporationFromAPI(corporationID)
	}
	return c, nil
}

func (m *Manager) getCorporationFromDB(corporationID int) (*Corporation, error) {
	c, err := m.pool.Open()
	if err != nil {
		return nil, err
	}
	defer m.pool.Release(c)
	r := c.QueryRow(
		`SELECT
			  c.corporation_id
			, c.name
			, c.alliance_id
			, c.creation_date
			, c.ticker
			, c.description
			FROM app.corporations c
			WHERE c.corporation_id = $1
				AND c.fetched_at > NOW() - INTERVAL '7 days'`, corporationID)
	char := &Corporation{}
	err = r.Scan(
		&char.CorporationID,
		&char.Name,
		&char.AllianceID,
		&char.CreationDate,
		&char.Ticker,
		&char.Description,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return char, nil
}

func (m *Manager) getCorporationFromAPI(corporationID int) (*Corporation, error) {
	char, err := m.eveapi.GetCorporation(corporationID)
	if err != nil {
		return nil, err
	}
	return m.apiCorporationToDB(char)
}

func (m *Manager) apiCorporationToDB(corp *eveapi.Corporation) (*Corporation, error) {
	db, err := m.pool.Open()
	if err != nil {
		return nil, err
	}
	defer m.pool.Release(db)
	c := &Corporation{
		CorporationID: corp.CorporationID,
		Name:          corp.Name,
		AllianceID:    corp.AllianceID,
		CreationDate:  corp.CreationDate,
		Ticker:        corp.Ticker,
		Description:   corp.Description,
	}
	_, err = db.Exec(
		`INSERT INTO app.corporations
			(corporation_id, name, alliance_id, creation_date, ticker, description)
			VALUES($1, $2, $3, $4, $5, $6)
			ON CONFLICT ON CONSTRAINT "corporations_pkey"
			DO UPDATE
			SET name = EXCLUDED.name
			  , alliance_id = EXCLUDED.alliance_id
			  , ticker = EXCLUDED.ticker
			  , description = EXCLUDED.description
			  , fetched_at = DEFAULT`,
		c.CorporationID,
		c.Name,
		c.AllianceID,
		c.CreationDate,
		c.Ticker,
		c.Description,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Divisions is a map of division key to division name.
type Divisions map[string]string

func (d Divisions) GetName(idx int) (string, bool) {
	id := strconv.Itoa(idx)
	if v, ok := d[id]; ok {
		return v, true
	}
	return "", false
}

func (d Divisions) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *Divisions) Scan(src interface{}) error {
	s, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("invalid value for division: %v", src)
	}
	return json.Unmarshal(s, &d)
}

type CorporationDetail struct {
	CorporationID int
	CEOID         int
	CEOName       string
	StationID     int
	StationName   string
	FactionID     int
	MemberCount   int
	Shares        int

	Wallets Divisions
	Hangars Divisions
}

func (m *Manager) GetCorporationDetail(corpID int) (*CorporationDetail, error) {
	d, err := m.getCorporationDetailFromDB(corpID)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, ErrCorpNotRegistered
	}
	return d, nil
}

func (m *Manager) FetchCorporationDetail(ctx context.Context) (*CorporationDetail, error) {
	return m.getCorporationDetailFromAPI(ctx)
}

func (m *Manager) getCorporationDetailFromDB(corporationID int) (*CorporationDetail, error) {
	c, err := m.pool.Open()
	if err != nil {
		return nil, err
	}
	defer m.pool.Release(c)
	r := c.QueryRow(
		`SELECT
			  c.corporation_id
			, c.ceo_id
			, c.ceo_name
			, c.hq_station_id
			, c.hq_station_name
			, c.faction_id
			, c.member_count
			, c.shares
			, c.hangars
			, c.divisions
			FROM app.corporation_details c
			WHERE c.corporation_id = $1`, corporationID)
	corp := &CorporationDetail{}
	err = r.Scan(
		&corp.CorporationID,
		&corp.CEOID,
		&corp.CEOName,
		&corp.StationID,
		&corp.StationName,
		&corp.FactionID,
		&corp.MemberCount,
		&corp.Shares,
		&corp.Hangars,
		&corp.Wallets,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return corp, nil
}

func (m *Manager) getCorporationDetailFromAPI(ctx context.Context) (*CorporationDetail, error) {
	sheet, err := m.eveapi.GetCorporationSheet(ctx)
	if err != nil {
		return nil, err
	}
	return m.apiCorporationDetailToDB(&CorporationDetail{
		CorporationID: sheet.CorporationID,
		CEOID:         sheet.CEOID,
		CEOName:       sheet.CEOName,
		StationID:     sheet.StationID,
		StationName:   sheet.StationName,
		FactionID:     sheet.FactionID,
		MemberCount:   sheet.MemberCount,
		Shares:        sheet.Shares,
		Hangars:       Divisions(sheet.Hangars),
		Wallets:       Divisions(sheet.Wallets),
	})
}

func (m *Manager) apiCorporationDetailToDB(detail *CorporationDetail) (*CorporationDetail, error) {
	db, err := m.pool.Open()
	if err != nil {
		return nil, err
	}
	defer m.pool.Release(db)
	_, err = db.Exec(
		`INSERT INTO app.corporation_details
			 (
			     corporation_id
			   , ceo_id
			   , ceo_name
			   , hq_station_id
			   , hq_station_name
			   , faction_id
			   , member_count
			   , shares
			   , hangars
			   , divisions
			 )
			 VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			 ON CONFLICT
			   ON CONSTRAINT "corporation_details_pkey"
			 DO UPDATE
			   SET ceo_id = EXCLUDED.ceo_id
			     , ceo_name = EXCLUDED.ceo_name
			     , hq_station_id = EXCLUDED.hq_station_id
			     , hq_station_name = EXCLUDED.hq_station_name
			     , faction_id = EXCLUDED.faction_id
			     , member_count = EXCLUDED.member_count
			     , shares = EXCLUDED.shares
			     , hangars = EXCLUDED.hangars
			     , divisions = EXCLUDED.divisions
			     , fetched_at = DEFAULT`,
		detail.CorporationID,
		detail.CEOID,
		detail.CEOName,
		detail.StationID,
		detail.StationName,
		detail.FactionID,
		detail.MemberCount,
		detail.Shares,
		detail.Hangars,
		detail.Wallets,
	)
	if err != nil {
		return nil, err
	}
	return detail, nil
}

type CorporationConfig struct {
	OptIn     bool
	OptInBy   int
	OptInDate time.Time

	CreatedBy int
	CreatedAt time.Time
}

func (m *Manager) GetCorporationsOptedIn() ([]int, error) {
	c, err := m.pool.Open()
	if err != nil {
		return nil, err
	}
	defer m.pool.Release(c)
	rs, err := c.Query(
		`SELECT
			  c.corporation_id
			FROM app.corporation_settings c
			WHERE c.opted_in = TRUE`)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var res []int
	for rs.Next() {
		i := 0
		if err = rs.Scan(&i); err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}

func (m *Manager) GetCorporationConfig(corpID int) (*CorporationConfig, error) {
	c, err := m.pool.Open()
	if err != nil {
		return nil, err
	}
	defer m.pool.Release(c)
	r := c.QueryRow(
		`SELECT
			  c.opted_in
			, c.opted_in_by
			, c.opted_in_at
			, c.created_by
			, c.created_at
			FROM app.corporation_settings c
			WHERE c.corporation_id = $1`, corpID)
	corp := &CorporationConfig{}
	err = r.Scan(
		&corp.OptIn,
		&corp.OptInBy,
		&corp.OptInDate,
		&corp.CreatedBy,
		&corp.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrCorpNotRegistered
		}
		return nil, err
	}
	return corp, nil
}

func (m *Manager) GetCorporationAuthorization(corpID int) (*Authorization, error) {
	config, err := m.GetCorporationConfig(corpID)
	if err != nil {
		return nil, err
	}
	if !config.OptIn {
		return nil, ErrCorpNotRegistered
	}
	user := &User{UserID: config.OptInBy}
	return m.GetAuthorization(user, RoleDirector)
}

func (m *Manager) SaveCorporationConfig(corpID int, detail *CorporationConfig) error {
	db, err := m.pool.Open()
	if err != nil {
		return err
	}
	defer m.pool.Release(db)
	_, err = db.Exec(
		`INSERT INTO app.corporation_settings
			 (
			     corporation_id
			   , opted_in
			   , opted_in_by
			   , opted_in_at
			   , created_by
			   , created_at
			 )
			 VALUES($1, $2, $3, $4, $5, DEFAULT)
			 ON CONFLICT
			   ON CONSTRAINT "corporation_settings_pkey"
			 DO UPDATE
			   SET opted_in = EXCLUDED.opted_in
			     , opted_in_by = EXCLUDED.opted_in_by
			     , opted_in_at = EXCLUDED.opted_in_at`,
		corpID,
		detail.OptIn,
		detail.OptInBy,
		detail.OptInDate,
		detail.CreatedBy,
	)
	return err
}

// UpdateCorporationData fetches updated data for all opted-in corporations.
//
// The function returned by this method is intended to be invoke in regular intervals.
func (m *Manager) UpdateCorporationDataFunc(logger log.Logger) func() error {
	return func() error {
		corps, err := m.GetCorporationsOptedIn()
		if err != nil {
			return err
		}
		if len(corps) == 0 {
			logger.Debugf("no corporations opted in, not updating corp data")
			return nil
		}
		for _, corpID := range corps {
			logger.Debugf("updating data for corp %d", corpID)
			a, err := m.GetCorporationAuthorization(corpID)
			if err != nil {
				logger.Errorf("error getting corp auth: %s", err.Error())
				continue
			}

			ctx := a.Context()
			if _, err := m.FetchCorporationDetail(ctx); err != nil {
				logger.Errorf("error fetching corp details: %s", err.Error())
			}
			if res, err := m.GetCorporationAssets(ctx, a.CorporationID); err != nil {
				logger.Errorf("error fetching corp assets: %s", err.Error())
			} else {
				logger.Debugf("fetched %d assets for corporation %d", len(res), a.CorporationID)
			}

			if res, err := m.GetCorporationOrders(ctx, a.CorporationID); err != nil {
				logger.Errorf("error fetching corp orders: %s", err.Error())
			} else {
				logger.Debugf("fetched %d orders for corporation %d", len(res), a.CorporationID)
			}

			if res, err := m.GetCorporationBlueprints(ctx, a.CorporationID); err != nil {
				logger.Errorf("error fetching corp blueprints: %s", err.Error())
			} else {
				logger.Debugf("fetched %d blueprints for corporation %d", len(res), a.CorporationID)
			}
		}
		return nil
	}
}