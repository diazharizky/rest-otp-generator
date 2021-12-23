package otp

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	otpKey = "some_otp_key"
)

var otpBase OTPBase

func init() {
	otpBase = OTPBase{
		Key:    otpKey,
		Period: 60 * time.Second,
		Digits: int8(4),
	}
}

func TestPasscodeLength(t *testing.T) {
	code, err := GenerateCode(otpBase)
	if err != nil {
		t.Errorf(err.Error())
	}

	codeLen := len(code)
	assert.Equal(t, codeLen, int(otpBase.Digits), fmt.Sprintf("OTP's length doesn't match, expected %d digits found %d.", otpBase.Digits, codeLen))
}

func TestValidVerification(t *testing.T) {
	code, err := GenerateCode(otpBase)
	require.NoError(t, err)

	valid, err := VerifyCode(OTPV{OTPBase: otpBase, Passcode: code})
	require.NoError(t, err)

	assert.Equal(t, true, valid, "OTP code verification shall be valid")
}

func TestInvalidVerification(t *testing.T) {
	valid, err := VerifyCode(OTPV{OTPBase: otpBase, Passcode: "0000"})
	require.NoError(t, err)

	assert.Equal(t, false, valid, "OTP code verification shall be invalid")
}
