package otp

import (
	"encoding/base32"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const (
	DigitsMin      = uint(4)
	DigitsMax      = uint(6)
	PeriodMin      = uint(60)
	PeriodMax      = uint(300)
	MaxAttemptsMin = uint(3)
	MaxAttemptsMax = uint(5)
)

type OTPBase struct {
	Key         string
	Period      uint `json:"period"`
	Digits      uint `json:"digits"`
	MaxAttempts uint `json:"max_attempts"`
	Attempts    uint `json:"attempts"`
}

type OTPV struct {
	OTPBase

	Passcode string `json:"passcode" validate:"required"`
}

func (p *OTPBase) FixParams() {
	if p.Digits < DigitsMin {
		p.Digits = DigitsMin
	}
	if p.Digits > DigitsMax {
		p.Digits = DigitsMax
	}
	if p.Period < PeriodMin {
		p.Period = PeriodMin
	}
	if p.Period > PeriodMax {
		p.Period = PeriodMax
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
