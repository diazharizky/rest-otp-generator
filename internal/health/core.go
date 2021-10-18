package health

import (
	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/diazharizky/rest-otp-generator/pkg/redis"
)

func init() {
	configs.Cfg.SetDefault("redis.host", "0.0.0.0")
	configs.Cfg.SetDefault("redis.port", 6379)
	configs.Cfg.SetDefault("redis.password", "")

	redisCfg := redis.Cfg{
		Host:     configs.Cfg.GetString("redis.host"),
		Port:     configs.Cfg.GetString("redis.port"),
		Password: configs.Cfg.GetString("redis.password"),
		Database: 0,
	}

	mCore.DB = &redis.Service{
		Client: redis.Connect(redisCfg),
	}
}

func (c *core) health() healthStatus {
	dbHealth := "OK"
	if err := c.DB.Health(); err != nil {
		dbHealth = "NOK"
	}

	return healthStatus{DB: dbHealth}
}
