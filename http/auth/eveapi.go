package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/antihax/goesi"
	"github.com/motki/motki/eveapi"
	"github.com/motki/motki/log"
	"github.com/motki/motki/model"
	"golang.org/x/oauth2"
)

const (
	eveAPIStateSessionKey = "__motki_eveapi_last_state"
	eveAPITokenSessionKey = "__motki_eveapi_token"
)

// EveAPIAuthorizer is an authentication authenticator using the EVE SSO.
type EveAPIAuthorizer struct {
	model  *model.Manager
	api    *eveapi.EveAPI
	logger log.Logger
}

// NewEveAPIAuthorizer creates a new authenticator using the given API client.
func NewEveAPIAuthorizer(m *model.Manager, api *eveapi.EveAPI, logger log.Logger) Authorizer {
	logger.Debugf("auth: init eveapi-based authorization")
	return &EveAPIAuthorizer{m, api, logger}
}

// BeginAuthorization starts the authorization process.
//
// The state is a nonce that is stored on the session and verified in the
// successful authentication request. This method redirects to the OAuth2
// endpoint on the EVE SSO site. Once the user logs in, they will be
// directed to the configured "return_url" which should call FinishAuth.
func (p *EveAPIAuthorizer) BeginAuthorization(s *Session, r model.Role, w http.ResponseWriter, req *http.Request) {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.RawURLEncoding.EncodeToString(b)
	s.Set(eveAPIStateSessionKey, state)
	s.Set(roleKey(state), int(r))
	url := p.api.AuthorizeURL(state, model.APIScopesForRole(r)...)
	http.Redirect(w, req, url, http.StatusFound)
}

// FinishAuthorization completes the authorization process.
//
// The state is loaded from the session and checked against the request.
// Once the user logs into EVE SSO, they should be directed to a handler
// that calls this method.
//
// The access token and refresh token, along with the character ID and name
// are stored on the session.
func (p *EveAPIAuthorizer) FinishAuthorization(s *Session, req *http.Request) error {
	state, ok := s.String(eveAPIStateSessionKey)
	s.Remove(eveAPIStateSessionKey)
	if !ok {
		return errors.New("didnt receive state")
	}
	if req.FormValue("state") != state {
		return errors.New("received state did not match expected")
	}
	k := roleKey(state)
	vr, ok := s.Int(k)
	s.Remove(k)
	if !ok {
		return errors.New("unable to find role for state")
	}
	r := model.Role(vr)
	token, err := p.api.TokenExchange(req.FormValue("code"))
	if err != nil {
		return err
	}
	source, err := p.api.TokenSource(token)
	if err != nil {
		return err
	}
	info, err := p.api.Verify(source)
	if err != nil {
		return err
	}
	// TODO: make a custom verification process for the given role
	b, err := json.Marshal(token)
	if err != nil {
		return err
	}
	if err = p.model.SaveAuthorization(s.user, r, int(info.CharacterID), (*oauth2.Token)(token)); err != nil {
		return err
	}
	s.Set(tokenKey(r), b)
	s.Set(characterIDKey, float64(info.CharacterID))
	return nil
}

func (p *EveAPIAuthorizer) InvalidateAuthorization(s *Session, r model.Role) error {
	s.Remove(tokenKey(r))
	s.Remove(characterIDKey)
	return nil
}

// AuthorizedContext creates a context with the authorized token information stored
// in an authenticated session.
//
// This method will attempt to load the necessary EVE API token for the given role from
// the user's authenticated session. If that fails, it will fall-back to loading the token
// from the database.
//
// Once loaded, the EVE API token is refreshed, updated in the database, and a context
// containing the token information is returned.
func (p *EveAPIAuthorizer) AuthorizedContext(s *Session, r model.Role) (context.Context, error) {
	source, ok := p.getTokenSourceFromSession(s, r)
	if !ok {
		source, ok = p.getTokenSourceFromDB(s, r)
		if !ok {
			return nil, errors.New("cannot get token information from session")
		}
	}
	info, err := p.api.Verify(source)
	if err != nil {
		return nil, err
	}
	t, err := source.Token()
	if err != nil {
		return nil, err
	}
	if err = p.model.SaveAuthorization(s.user, r, int(info.CharacterID), t); err != nil {
		return nil, err
	}
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	s.Set(tokenKey(r), b)
	s.Set(characterIDKey, float64(info.CharacterID))
	return context.WithValue(context.Background(), goesi.ContextOAuth2, source), nil
}

func (p *EveAPIAuthorizer) getTokenSourceFromDB(s *Session, r model.Role) (oauth2.TokenSource, bool) {
	a, err := p.model.GetAuthorization(s.user, r)
	if err != nil {
		p.logger.Debugf("unable to get authorization information from database: %s", err.Error())
		return nil, false
	}
	source, err := p.api.TokenSource(a.Token)
	if err != nil {
		p.logger.Warnf("unable to get valid TokenSource from database-loaded authorization: %s", err.Error())
		return nil, false
	}
	return source, true
}

func (p *EveAPIAuthorizer) getTokenSourceFromSession(s *Session, r model.Role) (oauth2.TokenSource, bool) {
	v, ok := s.Bytes(tokenKey(r))
	if !ok {
		return nil, false
	}
	tok := &oauth2.Token{}
	if err := json.Unmarshal(v, &tok); err != nil {
		return nil, false
	}
	source, err := p.api.TokenSource(tok)
	if err != nil {
		p.logger.Warnf("unable to get valid TokenSource from session-loaded authorization: %s", err.Error())
		return nil, false
	}
	return source, true
}

func roleKey(state string) string {
	return fmt.Sprintf("__motki_authing_%s", state)
}

func tokenKey(r model.Role) string {
	return fmt.Sprintf("%s_%v", eveAPITokenSessionKey, r)
}
