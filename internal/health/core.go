package health

import (
	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/diazharizky/rest-otp-generator/pkg/redis"
)

func New() {

}

func init() {
	configs.Cfg.SetDefault("redis.host", "0.0.0.0")
	configs.Cfg.SetDefault("redis.port", 6379)

	c.Redis = &redis.RDB{
		Client: redis.Connect(configs.Cfg.GetString("redis.host"), configs.Cfg.GetString("redis.port"), configs.Cfg.GetString("redis.password")),
	}
}

func (c *core) Health() HealthStatus {
	redisHealth := "OK"
	if err := c.Redis.Health(); err != nil {
		redisHealth = "NOK"
	}

	return HealthStatus{
		Redis: redisHealth,
	}
}
