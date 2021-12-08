package otp

import (
	"encoding/base32"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type OTPBase struct {
	Key      string
	Digits   int8          `json:"digits" validate:"required" redis:"digits"`
	Period   time.Duration `json:"period"`
	Attempts int8          `json:"attempts" redis:"attempts"`
}

type OTP struct {
	OTPBase

	Passcode string `json:"passcode" redis:"passcode"`
}

type OTPV struct {
	OTPBase

	Passcode string `json:"passcode" redis:"passcode" validate:"required"`
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
