package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RDB struct {
	Client *redis.Client
}

var ctx = context.Background()

func (r *RDB) Set(key string, value interface{}, expired time.Duration) error {
	err := r.Client.Set(ctx, key, value, expired).Err()
	return err
}

func (r *RDB) Get(key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}

	return val, nil
}

func (r *RDB) Remove(key string) error {
	err := r.Client.Del(ctx, key).Err()
	return err
}
