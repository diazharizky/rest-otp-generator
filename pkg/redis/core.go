package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	"github.com/go-redis/redis/v8"
)

type RDB struct {
	Client *redis.Client
}

func Connect(host string, port string, password string) *redis.Client {
	addr := fmt.Sprintf("%s:%s", host, port)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
}

func (r *RDB) Get(ctx context.Context, p otp.OTP) error {
	exists, err := r.Client.Exists(ctx, p.Key).Result()
	if err != nil {
		return err
	}

	if exists == 0 {
		return nil
	}

	if err = r.Client.HGetAll(ctx, p.Key).Scan(&p); err != nil {
		return err
	}

	return nil
}

func (r *RDB) Upsert(ctx context.Context, p otp.OTP) (err error) {
	fVal, err := toMSI(p)
	if err != nil {
		return err
	}

	if err = r.Client.HSet(ctx, p.Key, fVal).Err(); err != nil {
		return
	}

	err = r.Client.Expire(ctx, p.Key, p.Period*time.Second).Err()
	return
}

func (r *RDB) Delete(ctx context.Context, id string) error {
	err := r.Client.Del(ctx, id).Err()
	return err
}

func (r *RDB) Health() error {
	ctx := context.Background()
	_, err := r.Client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

// Convert any type of value to map[string]interface{}
func toMSI(val interface{}) (interface{}, error) {
	m, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}

	var u map[string]interface{}
	if err = json.Unmarshal(m, &u); err != nil {
		return nil, err
	}

	return u, nil
}
