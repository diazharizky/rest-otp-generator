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

func (r *RDB) HGetAll(key string) *redis.StringStringMapCmd {
	return r.Client.HGetAll(ctx, key)
}

func (r *RDB) HRemove(key string) error {
	err := r.Client.HDel(ctx, key).Err()
	return err
}
