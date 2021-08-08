package storages

import (
	"errors"
	"sync"
)

type MemoryStorage struct {
	table sync.Map
}

func (storage *MemoryStorage) Get(key string) (interface{}, error) {
	if storage == nil {
		return nil, errors.New("memory not configured")
	}
	value, ok := storage.table.Load(key)
	if !ok {
		return nil, errors.New("key not found")
	}
	return value, nil
}

func (storage *MemoryStorage) Set(key string, value interface{}) error {
	if storage == nil {
		return errors.New("memory not configured")
	}
	storage.table.Store(key, value)
	return nil
}

func NewMemoryStorage() *MemoryStorage {
	s := &MemoryStorage{
		table: sync.Map{},
	}
	return s
}
