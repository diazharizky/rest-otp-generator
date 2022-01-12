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

func (r *otpCache) Get(otpKey string, p *domain.OTP) (bool, error) {
	byt, err := cacheService.Get(r.client, otpKey)
	if err != nil {
		return false, err
	}
	if byt == nil {
		return false, nil
	}

	err = json.Unmarshal(byt, &p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *otpCache) Upsert(p domain.OTP) error {
	period := p.Period * uint(time.Second)
	return cacheService.Set(r.client, p.Key, p, time.Duration(period))
}

func (r *otpCache) Delete(otpKey string) error {
	return cacheService.Del(r.client, otpKey)
}
