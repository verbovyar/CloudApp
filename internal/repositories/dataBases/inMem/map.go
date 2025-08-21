package inMem

import (
	"errors"
	"sync"
)

var errorNoSuchKey = errors.New("no such key")

type Store struct {
	rep map[string]string

	mu sync.RWMutex
}

func New() *Store {
	return &Store{
		rep: make(map[string]string),
		mu:  sync.RWMutex{},
	}
}

func (s *Store) Put(key, value string) error {
	s.mu.Lock()
	s.rep[key] = value
	s.mu.Unlock()

	return nil
}

func (s *Store) Get(key string) (string, error) {
	s.mu.RLock()
	value, ok := s.rep[key]
	s.mu.RUnlock()
	if !ok {
		return "", errorNoSuchKey
	}

	return value, nil
}

func (s *Store) Delete(key string) error {
	s.mu.RLock()
	_, ok := s.rep[key]
	s.mu.RUnlock()
	if !ok {
		return errorNoSuchKey
	}

	delete(s.rep, key)

	return nil
}
