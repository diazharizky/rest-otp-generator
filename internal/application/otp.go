package application

import (
	"github.com/diazharizky/rest-otp-generator/internal/domain"
	"github.com/diazharizky/rest-otp-generator/internal/domain/repository"
	"github.com/diazharizky/rest-otp-generator/pkg/otp"
)

type otpApp struct {
	r repository.OTPRepository
}

type OTPAppInterface interface {
	GenerateOTP(*domain.OTP) (string, error)
	VerifyOTP(*domain.OTP) (bool, error)
}

func NewOTPApp(or repository.OTPRepository) otpApp {
	return otpApp{or}
}

func (i *otpApp) GenerateOTP(p *domain.OTP) (string, error) {
	code, err := otp.GenerateCode(p.Key, p.Period, int(p.Digits))
	if err != nil {
		return "", err
	}

	if err = i.r.Upsert(*p); err != nil {
		return "", err
	}
	return code, nil
}

func (i *otpApp) VerifyOTP(p *domain.OTP) (bool, error) {
	code := p.Code
	exists, err := i.r.Get(p.Key, p)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}

	if p.Attempts >= p.MaxAttempts {
		i.r.Delete(p.Key)
		return false, nil
	}

	valid, err := otp.VerifyCode(p.Key, code, p.Period, p.Digits)
	if err != nil {
		return false, err
	}
	if !valid {
		p.Attempts += 1
		if err = i.r.Upsert(*p); err != nil {
			return false, err
		}
		return false, nil
	}

	if err = i.r.Delete(p.Key); err != nil {
		return false, err
	}
	return true, nil
}
