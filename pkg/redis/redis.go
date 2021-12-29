package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/diazharizky/rest-otp-generator/pkg/db"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	db.Config

	Client *redis.Client
}

type handler struct {
	client *redis.Client
}

var Handler handler

func init() {
	host := os.Getenv("CACHE_HOST")
	if len(host) <= 0 {
		host = "0.0.0.0"
	}

	port := os.Getenv("CACHE_PORT")
	if len(port) <= 0 {
		port = "6379"
	}

	passwd := os.Getenv("CACHE_PASSWORD")
	if len(passwd) <= 0 {
		passwd = ""
	}

	dbName := os.Getenv("CACHE_DB")
	if len(dbName) <= 0 {
		dbName = "0"
	}

	configs.Cfg.Set("cache.host", host)
	configs.Cfg.Set("cache.port", port)
	configs.Cfg.Set("cache.password", passwd)
	configs.Cfg.Set("cache.db", dbName)

	dbNameInt, err := strconv.Atoi(dbName)
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	Handler = handler{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: passwd,
			DB:       dbNameInt,
		}),
	}
}

func (r *handler) Health() error {
	ctx := context.Background()
	return r.client.Ping(ctx).Err()
}

func (r *handler) Get(ctx context.Context, key string) (res []byte, err error) {
	exists, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if exists == 0 {
		return nil, nil
	}

	cmd := r.client.Get(ctx, key)
	return cmd.Bytes()
}

func (r *handler) Set(ctx context.Context, key string, value interface{}, duration time.Duration) (err error) {
	jbt, err := json.Marshal(value)
	if err != nil {
		return
	}
	err = r.client.Set(ctx, key, string(jbt), duration).Err()
	return
}

func (r *handler) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *handler) Close() error {
	return r.client.Close()
}
