package storages

import (
	"errors"
)

type MemoryStorage struct {
	table map[string]interface{}
}

func (storage *MemoryStorage) Get(key string) (interface{}, error) {
	if storage == nil {
		return nil, errors.New("memory not configured")
	}
	value, ok := storage.table[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return value, nil
}

func (storage *MemoryStorage) Set(key string, value interface{}) error {
	if storage == nil {
		return errors.New("memory not configured")
	}
	storage.table[key] = value
	return nil
}

func New() *MemoryStorage {
	s := &MemoryStorage{
		table: map[string]interface{}{"init": true},
	}
	return s
}
