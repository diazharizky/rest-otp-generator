package otp_test

import (
	"fmt"
	"testing"

	"github.com/diazharizky/rest-otp-generator/pkg/otp"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	otpKey = "some_otp_key"
)

var otpBase otp.OTPBase

func init() {
	otpBase = otp.OTPBase{
		Key:    otpKey,
		Period: 60,
		Digits: 4,
	}
}

func TestPasscodeLength(t *testing.T) {
	code, err := otp.GenerateCode(otpBase)
	if err != nil {
		t.Errorf(err.Error())
	}
	codeLen := len(code)
	assert.Equal(t, codeLen, int(otpBase.Digits), fmt.Sprintf("OTP's length doesn't match, expected %d digits found %d.", otpBase.Digits, codeLen))
}

func TestValidVerification(t *testing.T) {
	code, err := otp.GenerateCode(otpBase)
	require.NoError(t, err)

	valid, err := otp.VerifyCode(otp.OTPV{OTPBase: otpBase, Passcode: code})
	require.NoError(t, err)

	assert.Equal(t, true, valid, "OTP code verification shall be valid")
}

func TestInvalidVerification(t *testing.T) {
	valid, err := otp.VerifyCode(otp.OTPV{OTPBase: otpBase, Passcode: "0000"})
	require.NoError(t, err)

	assert.Equal(t, false, valid, "OTP code verification shall be invalid")
}

func TestSetDefaultValuesFuncMin(t *testing.T) {
	otpBase.Digits = 1
	otpBase.Period = 10
	otpBase.MaxAttempts = 1

	otpBase.FixParams()

	assert.Equal(t, otp.DigitsMin, otpBase.Digits)
	assert.Equal(t, otp.PeriodMin, otpBase.Period)
	assert.Equal(t, otp.MaxAttemptsMin, otpBase.MaxAttempts)
}

func TestSetDefaultValuesFuncMax(t *testing.T) {
	otpBase.Digits = 7
	otpBase.Period = 500
	otpBase.MaxAttempts = 7

	otpBase.FixParams()

	assert.Equal(t, otp.DigitsMax, otpBase.Digits)
	assert.Equal(t, otp.PeriodMax, otpBase.Period)
	assert.Equal(t, otp.MaxAttemptsMax, otpBase.MaxAttempts)
}
