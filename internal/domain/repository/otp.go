package repository

import "github.com/diazharizky/rest-otp-generator/internal/domain"

type OTPRepository interface {
	Get(otpKey string, param *domain.OTP) error
	Upsert(param domain.OTP) error
	Delete(otpKey string) error
}
