package otp

import (
	"encoding/base32"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateCode(key string, period uint, digits int) (string, error) {
	secret := base32.StdEncoding.EncodeToString([]byte(key))
	code, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period: period,
		Digits: otp.Digits(digits),
	})
	if err != nil {
		return "", err
	}
	return code, nil
}

func VerifyCode(key string, code string, period uint, digits uint) (bool, error) {
	secret := base32.StdEncoding.EncodeToString([]byte(key))
	valid, err := totp.ValidateCustom(code, secret, time.Now(), totp.ValidateOpts{
		Period: period,
		Digits: otp.Digits(digits),
	})
	if err != nil {
		return false, err
	}
	return valid, nil
}
