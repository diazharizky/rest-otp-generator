package redis

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
	Client   *redis.Client
}

type RedisHandler struct {
	client *redis.Client
}

func init() {
	redisHost := os.Getenv("REDIS_HOST")
	if len(redisHost) > 0 {
		configs.Cfg.Set("redis.host", redisHost)
	}
	redisPort := os.Getenv("REDIS_PORT")
	if len(redisPort) > 0 {
		configs.Cfg.Set("redis.port", redisPort)
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")
	if len(redisPassword) > 0 {
		configs.Cfg.Set("redis.password", redisPassword)
	}
	redisDB := os.Getenv("REDIS_DB")
	if len(redisDB) > 0 {
		configs.Cfg.Set("redis.db", redisDB)
	}
}

func GetHandler(client *redis.Client) *RedisHandler {
	return &RedisHandler{client: client}
}

func (r *RedisHandler) Health() (err error) {
	ctx := context.Background()
	_, err = r.client.Ping(ctx).Result()
	if err != nil {
		return
	}
	return
}

func (r *RedisHandler) Get(ctx context.Context, key string) (res []byte, err error) {
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

func (r *RedisHandler) Set(ctx context.Context, key string, value interface{}, duration time.Duration) (err error) {
	jbt, err := json.Marshal(value)
	if err != nil {
		return
	}
	err = r.client.Set(ctx, key, string(jbt), duration).Err()
	return
}

func (r *RedisHandler) Delete(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	return err
}
