package db

import (
	"context"

	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	"github.com/diazharizky/rest-otp-generator/pkg/redis"
)

type Database interface {
	Health() error

	Get(context.Context, *otp.OTP) error
	Upsert(context.Context, otp.OTP) error
	Delete(ctx context.Context, id string) error
}

func GetCfg() redis.Cfg {
	configs.Cfg.SetDefault("redis.host", "0.0.0.0")
	configs.Cfg.SetDefault("redis.port", 6379)
	configs.Cfg.SetDefault("redis.password", "")
	configs.Cfg.SetDefault("redis.db", 0)

	return redis.Cfg{
		Host:     configs.Cfg.GetString("redis.host"),
		Port:     configs.Cfg.GetString("redis.port"),
		Password: configs.Cfg.GetString("redis.password"),
		Database: configs.Cfg.GetInt("redis.db"),
	}
}
