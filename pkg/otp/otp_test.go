package otp_test

import (
	"fmt"
	"testing"

	"github.com/diazharizky/rest-otp-generator/pkg/otp"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	key    = "some_otp_key"
	period = 60
	digits = 4
	code   = "0000"
)

func TestPasscodeLength(t *testing.T) {
	code, err := otp.GenerateCode(key, period, digits)
	if err != nil {
		t.Errorf(err.Error())
	}
	codeLen := len(code)
	assert.Equal(t, codeLen, int(digits), fmt.Sprintf("OTP's length doesn't match, expected %d digits found %d.", digits, codeLen))
}

func TestValidVerification(t *testing.T) {
	code, err := otp.GenerateCode(key, period, digits)
	require.NoError(t, err)

	valid, err := otp.VerifyCode(key, code, period, digits)
	require.NoError(t, err)

	assert.Equal(t, true, valid, "OTP code verification shall be valid")
}

func TestInvalidVerification(t *testing.T) {
	valid, err := otp.VerifyCode(key, code, period, digits)
	require.NoError(t, err)

	assert.Equal(t, false, valid, "OTP code verification shall be invalid")
}
