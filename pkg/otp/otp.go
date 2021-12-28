package otp

import (
	"encoding/base32"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const (
	DigitsMin      = int8(4)
	DigitsMax      = int8(6)
	PeriodMin      = 60 * time.Second
	PeriodMax      = 300 * time.Second
	MaxAttemptsMin = int8(3)
	MaxAttemptsMax = int8(5)
)

type OTPBase struct {
	Key         string
	Period      time.Duration `json:"period"`
	Digits      int8          `json:"digits"`
	MaxAttempts int8          `json:"max_attempts"`
	Attempts    int8          `json:"attempts"`
}

type OTPV struct {
	OTPBase

	Passcode string `json:"passcode" validate:"required"`
}

func (p *OTPBase) SetDefaultValues() {
	if p.Digits < DigitsMin {
		p.Digits = DigitsMin
	}
	if p.Digits > DigitsMax {
		p.Digits = DigitsMax
	}
	if p.Period > PeriodMax {
		p.Period = PeriodMax
	}
	if p.Period < PeriodMin {
		p.Period = PeriodMin
	}
	if p.MaxAttempts < MaxAttemptsMin {
		p.MaxAttempts = MaxAttemptsMin
	}
	if p.MaxAttempts > MaxAttemptsMax {
		p.MaxAttempts = MaxAttemptsMax
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
