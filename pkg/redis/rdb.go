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

func (r *RDB) HSet(key string, value interface{}, expired time.Duration) (err error) {
	if err = r.Client.HSet(ctx, key, value).Err(); err != nil {
		return
	}

	err = r.Client.Expire(ctx, key, expired).Err()
	return
}

func (r *RDB) HGetAll(key string) (*redis.StringStringMapCmd, error) {
	exists, err := r.Client.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if exists == 0 {
		return nil, nil
	}
	return r.Client.HGetAll(ctx, key), nil
}

func (r *RDB) Remove(key string) error {
	err := r.Client.Del(ctx, key).Err()
	return err
}
