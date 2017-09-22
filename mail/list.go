package mail

import (
	"sync"

	"github.com/motki/motki/model"
)

type List interface {
	Add(Recipient) error
	Exists(Recipient) bool
}

type nilList struct{}

func (n nilList) Add(Recipient) error {
	return nil
}

func (n nilList) Exists(Recipient) bool {
	return false
}

type modelList struct {
	model *model.Manager

	key         string
	subscribers map[string]struct{}

	mutex *sync.RWMutex
}

func NewModelList(m *model.Manager, key string) (*modelList, error) {
	subs, err := m.GetMailingList(key)
	if err != nil {
		return nil, err
	}
	s := make(map[string]struct{})
	for _, r := range subs {
		s[r.Email] = struct{}{}
	}
	return &modelList{
		model:       m,
		key:         key,
		subscribers: s,
		mutex:       &sync.RWMutex{},
	}, nil
}

func (m *modelList) Add(rec Recipient) error {
	m.mutex.RLock()
	_, ok := m.subscribers[rec.Email]
	m.mutex.RUnlock()
	if ok {
		return nil
	}
	err := m.model.AddToMailingList(m.key, model.MailingListSubscriber{
		Name:  rec.Name,
		Email: rec.Email,
	})
	if err != nil {
		return err
	}
	m.mutex.Lock()
	m.subscribers[rec.Email] = struct{}{}
	m.mutex.Unlock()
	return nil
}

func (m *modelList) Exists(rec Recipient) bool {
	m.mutex.RLock()
	_, ok := m.subscribers[rec.Email]
	m.mutex.RUnlock()
	return ok
}
