package otp

import (
	"encoding/base32"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type OTPBase struct {
	Key         string
	Period      time.Duration `json:"period"`
	Digits      int8          `json:"digits"`
	MaxAttempts int8          `json:"max_attempts" redis:"max_attempts"`
	Attempts    int8          `json:"attempts" redis:"attempts"`
}

type OTPV struct {
	OTPBase
	Passcode string `json:"passcode" validate:"required"`
}

func (p *OTPBase) SetDefaultValues() {
	if p.Digits < 4 {
		p.Digits = 4
	}
	if p.Digits > 6 {
		p.Digits = 6
	}
	if p.Period > 300 {
		p.Period = 300 * time.Second
	}
	if p.Period < 60 {
		p.Period = 60 * time.Second
	}
}

func GenerateCode(p OTPBase) (string, error) {
	secret := base32.StdEncoding.EncodeToString([]byte(p.Key))
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period: uint(p.Period),
		Digits: otp.Digits(p.Digits),
	})
	if err != nil {
		return "", err
	}
	return passcode, nil
}

func VerifyCode(p OTPV) (bool, error) {
	secret := base32.StdEncoding.EncodeToString([]byte(p.Key))
	valid, err := totp.ValidateCustom(p.Passcode, secret, time.Now(), totp.ValidateOpts{
		Period: uint(p.Period),
		Digits: otp.Digits(p.Digits),
	})
	if err != nil {
		return false, err
	}
	return valid, nil
}
