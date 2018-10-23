package session

import (
	"net/http"
	"sync"

	"github.com/pkg/errors"
)

// A memoryManager handles persisting sessions to program memory.
type memoryManager struct {
	regenerator

	mut      sync.RWMutex
	sessions map[string]*Session

	cookie *cookier
}

func newMemoryManager(ckr *cookier) *memoryManager {
	m := &memoryManager{
		sessions: make(map[string]*Session),
		cookie:   ckr,
	}
	m.regenerator.m = m
	return m
}

func (m *memoryManager) Get(r *http.Request, w http.ResponseWriter) (*Session, error) {
	sess, err := m.loadSession(r)
	if err != nil {
		sess, err = newSession(m)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create new session")
		}
	}
	m.cookie.Set(sess.id, w)
	return sess, nil
}

func (m *memoryManager) Invalidate(r *http.Request, w http.ResponseWriter) {
	m.cookie.Remove(w)
	sess, err := m.loadSession(r)
	if err != nil {
		m.mut.Lock()
		defer m.mut.Unlock()
		delete(m.sessions, sess.id)
	}
}

// loadSession attempts to load an existing session.
func (m *memoryManager) loadSession(r *http.Request) (*Session, error) {
	sid, ok := m.cookie.Get(r)
	if !ok {
		return nil, errors.New("unable to get session ID from request")
	}
	m.mut.RLock()
	defer m.mut.RUnlock()
	if s, ok := m.sessions[sid]; ok {
		return s, nil
	}
	return nil, errors.New("session not found")
}

func (m *memoryManager) save(s *Session) error {
	m.mut.Lock()
	defer m.mut.Unlock()
	m.sessions[s.id] = s
	return nil
}

func (m *memoryManager) destroy(s *Session) error {
	m.mut.Lock()
	defer m.mut.Unlock()
	delete(m.sessions, s.id)
	return nil
}
