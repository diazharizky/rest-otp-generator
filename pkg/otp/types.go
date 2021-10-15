package otp

import (
	"time"
)

type BaseOTPPayload struct {
	Key    string
	Period time.Duration `json:"period" validate:"required"`
	Digits int8          `json:"digits" validate:"required"`
}

type VerifyOTPPayload struct {
	BaseOTPPayload

	Passcode string `json:"passcode" validate:"required"`
}
