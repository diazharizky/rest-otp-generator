package cache

import (
	"github.com/diazharizky/rest-otp-generator/internal/domain/repository"
	cacheService "github.com/diazharizky/rest-otp-generator/pkg/cache"
	"github.com/go-redis/redis/v8"
)

type healthCache struct {
	client *redis.Client
}

var _ repository.HealthRepository = &healthCache{}

func (h *healthCache) CacheHealth() error {
	return cacheService.Check(h.client)
}
