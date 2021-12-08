package redis

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	"github.com/go-redis/redis/v8"
)

const (
	messageInvalidOTP = "invalid OTP"
)

type Cfg struct {
	Host     string
	Port     string
	Password string
	Database int
}

type Service struct {
	Client *redis.Client
}

func Connect(cfg Cfg) *redis.Client {
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
