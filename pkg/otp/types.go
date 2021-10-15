package otp

import (
	"time"
)

type OTP struct {
	Key      string
	Period   time.Duration `json:"period" validate:"required"`
	Digits   int8          `json:"digits" validate:"required"`
	Passcode string        `json:"passcode" validate:"required" redis:"passcode"`
	Attempts int8          `json:"attempts" redis:"attempts"`
}
