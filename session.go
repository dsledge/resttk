package resttk

import (
	"crypto/rand"
	"fmt"
	"sync"
)

type SessionStoreInterface interface {
	Add(key string, value interface{})
	Get(key string, found bool)
	Delete(key string)
}

type SessionMemoryStore struct {
	sync.RWMutex
	M map[string]interface{}
}

func NewSessionMemoryStore() *SessionMemoryStore {
	sessionStore := &SessionMemoryStore{}
	sessionStore.M = make(map[string]interface{})
	return sessionStore
}

func (s *SessionMemoryStore) Add(key string, value interface{}) {
	s.Lock()
	defer s.Unlock()
	s.M[key] = value
}

func (s *SessionMemoryStore) Get(key string) (interface{}, bool) {
	s.Lock()
	defer s.Unlock()
	value, ok := s.M[key]
	return value, ok
}

func (s *SessionMemoryStore) Delete(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.M, key)
}

// Generate a session id
func GenerateSessionId() (string, error) {
	id := make([]byte, 32)
	if _, err := rand.Read(id); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", id), nil
}
