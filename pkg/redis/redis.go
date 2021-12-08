package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	"github.com/go-redis/redis/v8"
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

type OTP struct {
	Attempts int8          `redis:"attempts"`
	Digits   int8          `redis:"digits"`
	Passcode string        `redis:"passcode"`
	Period   time.Duration `redis:"period"`
}

func Connect(cfg Cfg) *redis.Client {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.Database,
	})
}

func (r *Service) Health() error {
	ctx := context.Background()
	_, err := r.Client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *Service) Get(ctx context.Context, p *otp.OTP) (err error) {
	exists, err := r.Client.Exists(ctx, p.Key).Result()
	if err != nil {
		return
	}

	if exists == 0 {
		return errors.New("invalid OTP")
	}

	var tmpOTP OTP
	if err = r.Client.HGetAll(ctx, p.Key).Scan(&tmpOTP); err != nil {
		return
	}

	p.Attempts = tmpOTP.Attempts
	p.Digits = tmpOTP.Digits
	p.Passcode = tmpOTP.Passcode
	p.Period = tmpOTP.Period

	return
}

func (r *Service) Upsert(ctx context.Context, p otp.OTP) (err error) {
	msi, err := toMSI(p)
	if err != nil {
		return
	}

	if err = r.Client.HSet(ctx, p.Key, msi).Err(); err != nil {
		return
	}

	err = r.Client.Expire(ctx, p.Key, p.Period*time.Second).Err()

	return
}

func (r *Service) Delete(ctx context.Context, id string) (err error) {
	err = r.Client.Del(ctx, id).Err()
	return
}

// Convert any type of value to map[string]interface{}
func toMSI(p interface{}) (interface{}, error) {
	m, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	var u map[string]interface{}
	if err = json.Unmarshal(m, &u); err != nil {
		return nil, err
	}

	return u, nil
}
