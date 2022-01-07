package cache

import (
	"encoding/json"
	"time"

	"github.com/diazharizky/rest-otp-generator/internal/domain"
	cacheService "github.com/diazharizky/rest-otp-generator/pkg/cache"
	"github.com/go-redis/redis/v8"
)

type otpCache struct {
	client *redis.Client
}

func (r *otpCache) Get(otpKey string, p *domain.OTP) error {
	byt, err := cacheService.Get(r.client, otpKey)
	if err != nil {
		return err
	}
	if byt == nil {
		return nil
	}
	err = json.Unmarshal(byt, &p)
	if err != nil {
		return err
	}
	return nil
}

func (r *otpCache) Upsert(p domain.OTP) error {
	return cacheService.Set(r.client, p.Key, p, time.Duration(p.Period))
}

func (r *otpCache) Delete(otpKey string) error {
	return cacheService.Del(r.client, otpKey)
}
