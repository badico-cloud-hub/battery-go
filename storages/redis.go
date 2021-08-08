package storages

import (
	"errors"
	"sync"
	"github.com/go-redis/redis/v8"
)

type MemoryStorage struct {
	table *redis.Clients
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

func New() *MemoryStorage {
	client := redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 0,
	 })
	 if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	 }

	 s := &MemoryStorage{
		table: s,
	}
	return s
}
