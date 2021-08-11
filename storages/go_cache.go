package storages

import (
	"errors"
	"log"

	"github.com/patrickmn/go-cache"
)

type GoCacheStorage struct {
	c *cache.Cache
}

func logGoCacheError(err error) {
	log.Printf("[GO CACHE ERROR]: %e", err)
}

var NotFoundErr error = errors.New("item not found")

func NewGoCacheStorage() *GoCacheStorage {
	c := cache.New(cache.NoExpiration, cache.NoExpiration)

	return &GoCacheStorage{
		c,
	}
}

func (gc *GoCacheStorage) Set(key string, value interface{}) error {
	gc.c.Set(key, value, cache.NoExpiration)

	return nil
}

func (gc *GoCacheStorage) Get(key string) (interface{}, error) {
	val, ok := gc.c.Get(key)

	if !ok {
		logGoCacheError(NotFoundErr)
		return "", NotFoundErr
	}

	return val, nil
}

func (gc *GoCacheStorage) Del(key string) (int64, error) {
	gc.c.Delete(key)

	return 0, nil
}
