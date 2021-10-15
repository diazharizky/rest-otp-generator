package otp

import (
	"context"
	"encoding/base32"
	"encoding/json"
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
	configs.Config.SetDefault("redis.host", "0.0.0.0")
	configs.Config.SetDefault("redis.port", 6379)

	c.DB.Client = redis.Connect(configs.Config.GetString("redis.host"), configs.Config.GetString("redis.port"), configs.Config.GetString("redis.password"))
}

func (o *core) generateOTP(p *otp.BaseOTPPayload) (string, error) {
	passcode, err := otp.GenerateOTP(p)
	if err != nil {
		return "", err
	}

	fVal, err := toMSI(storedOTP{Passcode: passcode, Attempts: 3})
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	if err = o.DB.HSetWithExp(ctx, p.Key, fVal, p.Period*time.Second); err != nil {
		return "", err
	}

	return passcode, nil
}

func (o *core) verifyOTP(p *otp.VerifyOTPPayload) error {
	ctx := context.Background()
	hOTP, err := o.DB.HGetAllWithCheck(ctx, p.Key)
	if err != nil {
		return err
	}

	if hOTP == nil {
		return errors.New("invalid OTP")
	}

	var sOTP storedOTP
	if err = hOTP.Scan(&sOTP); err != nil {
		return err
	}

	if sOTP.Attempts <= 0 {
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
		sOTP.Attempts -= 1
		fVal, err := toMSI(sOTP)
		if err != nil {
			return err
		}

		if err = o.DB.HSetWithExp(ctx, p.Key, fVal, p.Period*time.Second); err != nil {
			return err
		}

		return errors.New("invalid OTP")
	}

	if err = o.DB.Del(ctx, p.Key); err != nil {
		return err
	}

	return nil
}

// Convert any type of value to map[string]interface{}
func toMSI(val interface{}) (interface{}, error) {
	js, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}

	var i map[string]interface{}
	if err = json.Unmarshal(js, &i); err != nil {
		return nil, err
	}

	return i, nil
}
