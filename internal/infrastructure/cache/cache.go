package cache

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type caches struct {
	HealthCache healthCache
	OTPCache    otpCache
}

func NewCache(host string, port string, passwd string, db int) caches {
	addr := fmt.Sprintf("%s:%s", host, port)
	rClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	return caches{
		HealthCache: healthCache{rClient},
		OTPCache:    otpCache{rClient},
	}
}
