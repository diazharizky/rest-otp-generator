package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RDB struct {
	Client *redis.Client
}

func Connect(host string, port string, password string) *redis.Client {
	addr := fmt.Sprintf("%s:%s", host, port)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       0,
		Password: password,
	})
}

func (r *RDB) HSetWithExp(ctx context.Context, key string, value interface{}, expired time.Duration) (err error) {
	if err = r.Client.HSet(ctx, key, value).Err(); err != nil {
		return
	}

	err = r.Client.Expire(ctx, key, expired).Err()
	return
}

func (r *RDB) HGetAllWithCheck(ctx context.Context, key string) (*redis.StringStringMapCmd, error) {
	exists, err := r.Client.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if exists == 0 {
		return nil, nil
	}

	return r.Client.HGetAll(ctx, key), nil
}

func (r *RDB) Del(ctx context.Context, key string) error {
	err := r.Client.Del(ctx, key).Err()
	return err
}
