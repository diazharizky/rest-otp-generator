package otp

import (
	"context"
	"errors"

	"github.com/diazharizky/rest-otp-generator/internal/db"
	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	"github.com/diazharizky/rest-otp-generator/pkg/redis"
)

type core struct {
	Db db.Database
}

var c core

func init() {
	c.Db = &redis.Service{Client: redis.Connect(db.GetCfg())}
}

func (c *core) generateOTP(p *otp.OTP) (err error) {
	code, err := otp.GenerateCode(p.OTPBase)
	if err != nil {
		return
	}

	p.Passcode = code
	ctx := context.Background()
	if err = c.Db.Upsert(ctx, *p); err != nil {
		return
	}

	return
}

func (c *core) verifyOTP(p otp.OTPV) (err error) {
	ctx := context.Background()
	o := otp.OTP{OTPBase: p.OTPBase}
	if err = c.Db.Get(ctx, &o); err != nil {
		return
	}

	p.OTPBase = o.OTPBase
	if p.Passcode == "" {
		return errors.New("invalid OTP")
	}

	if p.Attempts <= 0 {
		return errors.New("invalid OTP")
	}

	valid, err := otp.VerifyCode(p)
	if err != nil {
		return
	}

	if !valid {
		o.Attempts -= 1
		if err = c.Db.Upsert(ctx, o); err != nil {
			return
		}

		return errors.New("invalid OTP")
	}

	if err = c.Db.Delete(ctx, p.Key); err != nil {
		return
	}

	return nil
}
