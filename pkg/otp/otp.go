package otp

import (
	"encoding/base32"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const (
	digitsMin      = 4
	digitsMax      = 6
	periodMin      = 60
	periodMax      = 300
	maxAttemptsMin = 3
	maxAttemptsMax = 5
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
	if p.Digits < digitsMin {
		p.Digits = digitsMin
	}
	if p.Digits > digitsMax {
		p.Digits = digitsMax
	}
	if p.Period > periodMax {
		p.Period = periodMax * time.Second
	}
	if p.Period < periodMin {
		p.Period = periodMin * time.Second
	}
	if p.MaxAttempts < maxAttemptsMin {
		p.MaxAttempts = maxAttemptsMin
	}
	if p.MaxAttempts > maxAttemptsMax {
		p.MaxAttempts = maxAttemptsMax
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
