// Package session provides session management capabilities.
package session

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/motki/motki/log"
	"github.com/pkg/errors"
)

const cookieName = "__motkid"

// Config is the configuration for a session manager.
type Config struct {
	Type      string `toml:"store"`
	HTTPSOnly bool   `toml:"secure_only"`
	Secret    string `toml:"secret"`

	FileConfig `toml:"file"`
}

// FileConfig represents a configuration for file-based sessions.
type FileConfig struct {
	StoragePath string `toml:"storage_path"`
}

// A Manager handles loading and saving of sessions.
type Manager interface {
	// Get loads or creates a session for the given request.
	Get(*http.Request, http.ResponseWriter) (*Session, error)

	// Invalidate destroys any session related to the given request, if any.
	Invalidate(*http.Request, http.ResponseWriter)

	// save saves the session to the underlying persistent storage.
	save(*Session) error
}

// A cookier is responsible for setting, getting, and removing session ID cookies.
type cookier struct {
	secureOnly bool
	secret     []byte
}

// Set writes a Set-Cookie header for the given session ID.
func (c *cookier) Set(sessID string, w http.ResponseWriter) {
	sig := c.sign(sessID)
	http.SetCookie(w, &http.Cookie{
		Name:    cookieName,
		Value:   sessID + ":" + sig,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24),
		Secure:  c.secureOnly,
	})
}

// Get attempts to get the session ID stored in the cookies in the given request.
func (c *cookier) Get(r *http.Request) (string, bool) {
	cook, err := r.Cookie(cookieName)
	if err != nil {
		return "", false
	}
	parts := strings.Split(cook.Value, ":")
	if len(parts) != 2 {
		return "", false
	}
	if !c.verify(parts[0], parts[1]) {
		return "", false
	}
	return parts[0], true
}

// Remove writes a Set-Cookie header to remove any session cookie.
func (c *cookier) Remove(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})
}

func (c *cookier) sign(message string) string {
	mac := hmac.New(sha256.New, c.secret)
	mac.Write([]byte(message))
	sum := mac.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(sum)
}

func (c *cookier) verify(message, signature string) bool {
	suspect, err := base64.RawURLEncoding.DecodeString(signature)
	if err != nil {
		// TODO: Pretty sure short-circuiting here is considered a timing leak.
		return false
	}
	mac := hmac.New(sha256.New, c.secret)
	mac.Write([]byte(message))
	expected := mac.Sum(nil)
	return hmac.Equal(suspect, expected)
}

// NewManager creates a new session manager with the given configuration.
func NewManager(c Config, l log.Logger) Manager {
	if string(c.Secret) == "" {
		l.Fatalf("session: secret cannot be empty")
	}
	var ret Manager
	cookie := &cookier{c.HTTPSOnly, []byte(c.Secret)}
	if c.HTTPSOnly {
		l.Debugf("session: secure only session cookies enabled; only https clients will have sessions")
	}
	if c.Type == "file" {
		l.Debugf("session: using file-based sessions")
		l.Debugf("session: storage path: %s", c.StoragePath)
		ret = newFileManager(c.StoragePath, cookie)
	} else {
		l.Debugf("session: using in-memory sessions")
		ret = newMemoryManager(cookie)
	}
	return ret
}

// A Session is a semi-persistent user storage container.
//
// Sessions are created by the Session middleware. Each Session is
// expected to live for only the associated Request'c lifetime and
// are not designed to be operated on concurrently.
type Session struct {
	m Manager

	id   string
	vals map[string]interface{}
}

// Flush saves the session to persistent storage.
func (s *Session) Flush() error {
	return s.m.save(s)
}

// SetFlash adds a message to the session with the given key.
//
// Flash messages are messages intended to be shown on the next page
// load. Reading a flash removes it from the session.
func (s *Session) SetFlash(key string, message string) {
	s.vals[key] = message
}

// Flash returns the stored flash message with the given key or an empty string.
//
// Calling this method removes the value stored in the session.
func (s *Session) Flash(key string) string {
	if v, ok := s.vals[key].(string); ok {
		s.Remove(key)
		return v
	}
	return ""
}

// NewCSRF creates a one-time use random token to be used as a CSRF token.
//
// Each form should contain a CSRF token that is generated when the form is shown
// and validated when a form is submitted.
func (s *Session) NewCSRF(key string) string {
	b := make([]byte, 16)
	rand.Read(b)
	tok := base64.RawURLEncoding.EncodeToString(b)
	s.Set(key, tok)
	return tok
}

// CheckCSRF checks the given suspect token against the stored CSRF token.
//
// Calling this method removes the value, if any, from the session.
func (s *Session) CheckCSRF(key string, suspect string) bool {
	e, ok := s.String(key)
	if !ok {
		return false
	}
	defer s.Remove(key)
	return e == suspect
}

// Set sets the given key to the given value.
func (s *Session) Set(key string, value interface{}) {
	s.vals[key] = value
}

// Get returns the value for the given string.
func (s *Session) Get(key string) (interface{}, bool) {
	v, ok := s.vals[key]
	return v, ok
}

// Bytes attempts to retrieve the value for the given key as a byte slice.
func (s *Session) Bytes(key string) ([]byte, bool) {
	v, ok := s.vals[key]
	if ok {
		switch s := v.(type) {
		case string:
			return []byte(s), true

		case []byte:
			return s, true
		}
	}
	return []byte{}, false
}

// String attempts to retrieve the value for the given key as a string.
func (s *Session) String(key string) (string, bool) {
	v, ok := s.vals[key]
	if ok {
		if s, ok := v.(string); ok {
			return s, true
		}
	}
	return "", false
}

// Int attempts to retrieve the value for the given key as an int.
func (s *Session) Int(key string) (int, bool) {
	v, ok := s.vals[key]
	if ok {
		if s, ok := v.(int); ok {
			return s, true
		} else if f, ok := v.(float64); ok {
			return int(f), true
		}
	}
	return 0, false
}

// Remove removes the given key from the session.
func (s *Session) Remove(key string) bool {
	delete(s.vals, key)
	return true
}

// newSession creates a new session.
func newSession(m Manager) (*Session, error) {
	id, err := newSessionID()
	if err != nil {
		return nil, errors.Wrap(err, "unable to generate session ID")
	}
	sess := &Session{
		m:    m,
		id:   id,
		vals: make(map[string]interface{}),
	}
	return sess, nil
}

// newSessionID creates a unique ID for use with a new session.
func newSessionID() (string, error) {
	b := make([]byte, 24)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
