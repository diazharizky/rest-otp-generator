package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

func Check(client *redis.Client) error {
	return client.Ping(context.Background()).Err()
}

func Get(client *redis.Client, key string) ([]byte, error) {
	ctx := context.Background()
	exists, err := client.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if exists == 0 {
		return nil, nil
	}
	return client.Get(ctx, key).Bytes()
}

func Set(client *redis.Client, key string, data interface{}, dur time.Duration) error {
	byt, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return client.Set(context.Background(), key, string(byt), dur).Err()
}

func Del(client *redis.Client, key string) error {
	return client.Del(context.Background(), key).Err()
}
