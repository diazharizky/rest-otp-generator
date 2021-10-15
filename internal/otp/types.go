package otp

import (
	"github.com/diazharizky/rest-otp-generator/pkg/redis"
)

type storedOTP struct {
	Passcode string `json:"passcode" redis:"passcode"`
	Attempts int8   `json:"attempts" redis:"attempts"`
}

type core struct {
	DB redis.RDB
}
