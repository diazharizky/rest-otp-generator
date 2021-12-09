package otp

import (
	"fmt"
	"testing"
	"time"
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
	if codeLen != int(otpBase.Digits) {
		t.Errorf(fmt.Sprintf("OTP's length doesn't match, expected %d digits found %d.", otpBase.Digits, codeLen))
	}
}

func TestVerification(t *testing.T) {
	code, err := GenerateCode(otpBase)
	if err != nil {
		t.Errorf(err.Error())
	}

	valid, err := VerifyCode(OTPV{OTPBase: otpBase, Passcode: code})
	if err != nil {
		t.Errorf(err.Error())
	}

	if !valid {
		t.Errorf(fmt.Sprintf("OTP's code doesn't valid, the code is %s.", code))
	}
}
