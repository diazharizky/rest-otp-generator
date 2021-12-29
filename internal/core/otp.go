package core

import (
	"context"
	"encoding/json"
	"time"

	"github.com/diazharizky/rest-otp-generator/pkg/otp"
)

func (c *core) GenerateOTP(p *otp.OTPBase) (code string, err error) {
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

func (c *core) VerifyOTP(p *otp.OTPV) (bool, error) {
	ctx := context.Background()
	jbt, err := c.DB.Get(ctx, p.Key)
	if err != nil {
		return false, err
	}
	if jbt == nil {
		return false, nil
	}

	err = json.Unmarshal(jbt, &p.OTPBase)
	if err != nil {
		return false, err
	}
	if p.Attempts >= p.MaxAttempts {
		if err = c.DB.Delete(ctx, p.Key); err != nil {
			return false, err
		}
		return false, nil
	}

	valid, err := otp.VerifyCode(*p)
	if err != nil {
		return false, err
	}
	if !valid {
		p.Attempts += 1
		if err = c.DB.Set(ctx, p.Key, p.OTPBase, time.Duration(p.Period)*time.Second); err != nil {
			return false, err
		}
		return false, nil
	}

	if err = c.DB.Delete(ctx, p.Key); err != nil {
		return false, err
	}
	return true, nil
}
