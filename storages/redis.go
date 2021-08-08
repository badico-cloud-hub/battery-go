package storages

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"context"
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

func NewRedisStorage() (*RedisStorage, error) {
	var (
		// ErrNil = errors.New("no matching record found in redis database")
		Ctx    = context.TODO()
	)
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "sOmE_sEcUrE_pAsS",
		DB: 0,
	 })
	 if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	 }

	 s := &RedisStorage{
		table: client,
	}
	return s, nil
}
