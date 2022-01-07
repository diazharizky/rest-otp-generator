package cache

import (
	cacheService "github.com/diazharizky/rest-otp-generator/pkg/cache"
	"github.com/go-redis/redis/v8"
)

type healthCache struct {
	client *redis.Client
}

func (h *healthCache) CacheHealth() error {
	return cacheService.Check(h.client)
}
