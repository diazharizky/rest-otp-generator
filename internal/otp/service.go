package otp

import (
	"context"
	"encoding/base32"
	"errors"
	"time"

	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	"github.com/diazharizky/rest-otp-generator/pkg/redis"
	otplib "github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

var c core

func init() {
	configs.Cfg.SetDefault("redis.host", "0.0.0.0")
	configs.Cfg.SetDefault("redis.port", 6379)

	c.DB = &redis.RDB{
		Client: redis.Connect(configs.Cfg.GetString("redis.host"), configs.Cfg.GetString("redis.port"), configs.Cfg.GetString("redis.password")),
	}
}

func (o *core) generateOTP(p otp.OTP) (string, error) {
	passcode, err := otp.GenerateOTP(p)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	if err = o.DB.Upsert(ctx, p); err != nil {
		return "", err
	}

	return passcode, nil
}

func (o *core) verifyOTP(p otp.OTP) error {
	ctx := context.Background()
	err := o.DB.Get(ctx, p)
	if err != nil {
		return err
	}

	if p.Attempts <= 0 {
		return errors.New("invalid OTP")
	}

	secret := base32.StdEncoding.EncodeToString([]byte(p.Key))
	valid, err := totp.ValidateCustom(p.Passcode, secret, time.Now(), totp.ValidateOpts{
		Period: uint(p.Period),
		Digits: otplib.Digits(p.Digits),
	})
	if err != nil {
		return err
	}

	if !valid {
		p.Attempts -= 1
		if err = o.DB.Upsert(ctx, p); err != nil {
			return err
		}

		return errors.New("invalid OTP")
	}

	if err = o.DB.Delete(ctx, p.Key); err != nil {
		return err
	}

	return nil
}
