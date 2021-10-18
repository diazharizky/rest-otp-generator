package otp

import (
	"context"
	"errors"

	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	"github.com/diazharizky/rest-otp-generator/pkg/redis"
)

var mCore core

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

func (c *core) generateOTP(p otp.OTP) error {
	code, err := otp.GenerateCode(p)
	if err != nil {
		return err
	}

	p.Passcode = code
	ctx := context.Background()
	if err = c.DB.Upsert(ctx, p); err != nil {
		return err
	}

	return nil
}

func (c *core) verifyOTP(p otp.OTP) error {
	ctx := context.Background()
	if err := c.DB.Get(ctx, p); err != nil {
		return err
	}

	if p.Passcode == "" {
		return errors.New("invalid OTP")
	}

	if p.Attempts <= 0 {
		return errors.New("invalid OTP")
	}

	valid, err := otp.VerifyCode(p)
	if err != nil {
		return err
	}

	if !valid {
		p.Attempts -= 1
		if err = c.DB.Upsert(ctx, p); err != nil {
			return err
		}

		return errors.New("invalid OTP")
	}

	if err = c.DB.Delete(ctx, p.Key); err != nil {
		return err
	}

	return nil
}
