package otp

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/diazharizky/rest-otp-generator/internal/db"
	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	cache "github.com/diazharizky/rest-otp-generator/pkg/redis"
)

const (
	messageInvalidOTP = "invalid OTP"
)

type core struct {
	DB db.Database
}

var c core

func init() {
	c.DB = &cache.Handler
}

func (c *core) generateOTP(p *otp.OTPBase) (code string, err error) {
	code, err = otp.GenerateCode(*p)
	if err != nil {
		return
	}

	ctx := context.Background()
	if err = c.DB.Set(ctx, p.Key, *p, time.Duration(p.Period)*time.Second); err != nil {
		return
	}
	return
}

func (c *core) verifyOTP(p *otp.OTPV) (err error) {
	ctx := context.Background()
	jbt, err := c.DB.Get(ctx, p.Key)
	if err != nil {
		return
	}
	if jbt == nil {
		return errors.New(messageInvalidOTP)
	}

	err = json.Unmarshal(jbt, &p.OTPBase)
	if err != nil {
		return
	}
	if p.Attempts >= p.MaxAttempts {
		if err = c.DB.Delete(ctx, p.Key); err != nil {
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
		if err = c.DB.Set(ctx, p.Key, p.OTPBase, time.Duration(p.Period)*time.Second); err != nil {
			return
		}
		return errors.New(messageInvalidOTP)
	}

	if err = c.DB.Delete(ctx, p.Key); err != nil {
		return
	}
	return nil
}
