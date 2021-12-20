package redis

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/diazharizky/rest-otp-generator/internal/db"
	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	"github.com/go-redis/redis/v8"
)

const (
	defaultHost     = "0.0.0.0"
	defaultPort     = "6379"
	defaultPassword = ""
	defaultDB       = 0

	messageInvalidOTP = "invalid OTP"
)

func GetCfg() db.Cfg {
	configs.Cfg.SetDefault("redis.host", defaultHost)
	configs.Cfg.SetDefault("redis.port", defaultPort)
	configs.Cfg.SetDefault("redis.password", defaultPassword)
	configs.Cfg.SetDefault("redis.db", defaultDB)

	host := os.Getenv("REDIS_HOST")
	if len(host) <= 0 {
		host = configs.Cfg.GetString("redis.host")
	}

	port := os.Getenv("REDIS_PORT")
	if len(port) <= 0 {
		port = configs.Cfg.GetString("redis.port")
	}

	password := os.Getenv("REDIS_PASSWORD")
	if len(password) <= 0 {
		password = configs.Cfg.GetString("redis.password")
	}

	dbIndex := os.Getenv("REDIS_DB")
	if len(dbIndex) <= 0 {
		dbIndex = configs.Cfg.GetString("redis.db")
	}

	dbInt, err := strconv.Atoi(dbIndex)
	if err != nil {
		panic(err)
	}

	return db.Cfg{
		Host:     host,
		Port:     port,
		Password: password,
		Database: dbInt,
	}
}

type Service struct {
	Client *redis.Client
}

func Connect(cfg db.Cfg) *redis.Client {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.Database,
	})
}

func (r *Service) Health() (err error) {
	ctx := context.Background()
	_, err = r.Client.Ping(ctx).Result()
	if err != nil {
		return
	}
	return
}

func (r *Service) Get(ctx context.Context, p *otp.OTPBase) (err error) {
	exists, err := r.Client.Exists(ctx, p.Key).Result()
	if err != nil {
		return
	}
	if exists == 0 {
		return errors.New(messageInvalidOTP)
	}
	if err = r.Client.HGetAll(ctx, p.Key).Scan(p); err != nil {
		return
	}
	return
}

func (r *Service) Upsert(ctx context.Context, p otp.OTPBase) (err error) {
	maxAttempts := strconv.Itoa(int(p.MaxAttempts))
	attempts := strconv.Itoa(int(p.Attempts))
	if err = r.Client.HSet(ctx, p.Key, []string{"max_attempts", maxAttempts, "attempts", attempts}).Err(); err != nil {
		return
	}
	err = r.Client.Expire(ctx, p.Key, p.Period).Err()
	return
}

func (r *Service) Delete(ctx context.Context, key string) (err error) {
	err = r.Client.Del(ctx, key).Err()
	return
}
