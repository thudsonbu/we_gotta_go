package crdt

import (
	"sync"
	"time"
)

type KeyValue struct {
	Value     interface{}
	Timestamp time.Time
}

type LWWKeyValueStore struct {
	store map[string]KeyValue
	lock  sync.RWMutex
}

func NewLWWKeyValueStore() *LWWKeyValueStore {
	return &LWWKeyValueStore{
		store: make(map[string]KeyValue),
	}
}

func (s *LWWKeyValueStore) Set(key string, value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	kv, exists := s.store[key]
	if !exists || kv.Timestamp.Before(time.Now()) {
		s.store[key] = KeyValue{
			Value:     value,
			Timestamp: time.Now(),
		}
	}
}

func (s *LWWKeyValueStore) Get(key string) (interface{}, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	kv, exists := s.store[key]
	if exists {
		return kv.Value, true
	}
	return nil, false
}

func (s *LWWKeyValueStore) Merge(other *LWWKeyValueStore) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for key, kv := range other.store {
		existingKV, exists := s.store[key]
		if !exists || existingKV.Timestamp.Before(kv.Timestamp) {
			s.store[key] = kv
		}
	}
}
