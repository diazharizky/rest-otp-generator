package domain_test

import (
	"testing"

	"github.com/diazharizky/rest-otp-generator/internal/domain"

	"github.com/stretchr/testify/assert"
)

var defOTP domain.OTP

func init() {
	defOTP = domain.OTP{}
}

func TestFixPropsMin(t *testing.T) {
	o := domain.OTP{Period: 40, Digits: 3, MaxAttempts: 2}
	o.FixProps()

	assert.Equal(t, domain.MinDigits, o.Digits)
	assert.Equal(t, domain.MinPeriod, o.Period)
	assert.Equal(t, domain.MinMaxAttempts, o.MaxAttempts)
}

func TestFixPropsMax(t *testing.T) {
	o := domain.OTP{Period: 350, Digits: 7, MaxAttempts: 6}
	o.FixProps()

	assert.Equal(t, domain.MaxDigits, o.Digits)
	assert.Equal(t, domain.MaxPeriod, o.Period)
	assert.Equal(t, domain.MaxMaxAttemtps, o.MaxAttempts)
}

func TestValidatePropsNoOTP(t *testing.T) {
	err := defOTP.ValidateProps()[0]
	assert.Equal(t, domain.OTPCodeRequiredMsg, err)
}

func TestValidatePropsOTPLenMismatch(t *testing.T) {
	defOTP.Code = "123"
	err := defOTP.ValidateProps()[0]
	assert.Equal(t, domain.MismatchOTPLenMsg, err)
}
