package session

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// A fileManager handles persisting session to the filesystem.
type fileManager struct {
	regenerator
	storagePath string

	cookie *cookier
}

func newFileManager(storagePath string, ckr *cookier) *fileManager {
	m := &fileManager{
		storagePath: storagePath,
		cookie:      ckr,
	}
	m.regenerator.m = m
	return m
}

func (m *fileManager) Get(r *http.Request, w http.ResponseWriter) (*Session, error) {
	sess, err := m.loadSession(r)
	if err != nil {
		// Couldn't load existing session, create a new one.
		sess, err = newSession(m)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create new session")
		}
	}
	m.cookie.Set(sess.id, w)
	return sess, nil
}

func (m *fileManager) Invalidate(r *http.Request, w http.ResponseWriter) {
	m.cookie.Remove(w)
	sess, err := m.loadSession(r)
	if err == nil {
		os.Remove(filepath.Join(m.storagePath, sess.id))
	}
}

// loadSession attempts to load an existing session.
func (m *fileManager) loadSession(r *http.Request) (*Session, error) {
	sid, ok := m.cookie.Get(r)
	if !ok {
		return nil, errors.New("unable to get session ID from request")
	}
	content, err := ioutil.ReadFile(filepath.Join(m.storagePath, sid))
	if err != nil {
		return nil, errors.Wrap(err, "unable to read session data from filesystem")
	}
	vals := make(map[string]interface{})
	if err = json.Unmarshal(content, &vals); err != nil {
		return nil, err
	}
	return &Session{
		m:    m,
		id:   sid,
		vals: vals,
	}, nil
}

func (m *fileManager) save(s *Session) error {
	p := filepath.Join(m.storagePath, s.id)
	b, err := json.Marshal(s.vals)
	if err != nil {
		return errors.Wrap(err, "unable to marshal session data to json")
	}
	return errors.Wrap(ioutil.WriteFile(p, b, 0600), "unable to save session to filesystem")
}

func (m *fileManager) destroy(s *Session) error {
	p := filepath.Join(m.storagePath, s.id)
	return os.Remove(p)
}
