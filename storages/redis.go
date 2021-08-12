package storages

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

type RedisStorage struct {
	table *redis.Client
}

func (storage *RedisStorage) Get(key string) (interface{}, error) {
	ctx := context.TODO()
	if storage == nil {
		return nil, errors.New("memory not configured")
	}
	value := storage.table.Get(ctx, key)
	return value.Val(), nil
}

func (storage *RedisStorage) Set(key string, value interface{}) error {
	ctx := context.TODO()
	if storage == nil {
		return errors.New("redis not configured")
	}
	storage.table.Set(ctx, key, value, 0)
	return nil
}

func NewRedisStorage(Addr,
	Password string,
	DB int) (*RedisStorage, error) {
	var (
		// ErrNil = errors.New("no matching record found in redis database")
		Ctx = context.TODO()
	)
	client := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       DB,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}

	s := &RedisStorage{
		table: client,
	}
	return s, nil
}
