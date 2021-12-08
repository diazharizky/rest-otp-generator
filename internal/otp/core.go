package otp

import (
	"context"
	"errors"

	"github.com/diazharizky/rest-otp-generator/internal/db"
	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	"github.com/diazharizky/rest-otp-generator/pkg/redis"
)

const (
	messageInvalidOTP = "invalid OTP"
)

type core struct {
	Db db.Database
}

var c core

func init() {
	client := redis.Connect(db.GetCfg())
	c.Db = &redis.Service{Client: client}
}

func (c *core) generateOTP(p *otp.OTPBase) (code string, err error) {
	code, err = otp.GenerateCode(*p)
	if err != nil {
		return
	}

	ctx := context.Background()
	if err = c.Db.Upsert(ctx, *p); err != nil {
		return
	}

	return
}

func (c *core) verifyOTP(p *otp.OTPV) (err error) {
	ctx := context.Background()
	if err = c.Db.Get(ctx, &p.OTPBase); err != nil {
		return
	}

	if p.Attempts >= p.MaxAttempts {
		if err = c.Db.Delete(ctx, p.Key); err != nil {
			return
		}
		return errors.New(messageInvalidOTP)
	}

	valid, err := otp.VerifyCode(*p)
	if err != nil {
		return
	}
	if !valid {
		p.Attempts += 1
		if err = c.Db.Upsert(ctx, p.OTPBase); err != nil {
			return
		}
		return errors.New(messageInvalidOTP)
	}

	if err = c.Db.Delete(ctx, p.Key); err != nil {
		return
	}

	return nil
}
