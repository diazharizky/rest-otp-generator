package cache

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
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

func CacheConfig() (string, string, string, int) {
	host := os.Getenv("CACHE_HOST")
	if len(host) <= 0 {
		host = "0.0.0.0"
	}

	port := os.Getenv("CACHE_PORT")
	if len(port) <= 0 {
		port = "6379"
	}

	passwd := os.Getenv("CACHE_PASSWORD")
	dbString := os.Getenv("CACHE_DB")
	if len(dbString) <= 0 {
		dbString = "0"
	}
	db, err := strconv.Atoi(dbString)
	if err != nil {
		panic(err)
	}

	return host, port, passwd, db
}
