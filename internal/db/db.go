package db

import (
	"context"

	"github.com/diazharizky/rest-otp-generator/pkg/otp"
)

type Cfg struct {
	Host     string
	Port     string
	Password string
	Database int
}

type Database interface {
	Health() error

	Get(context.Context, *otp.OTPBase) error
	Upsert(context.Context, otp.OTPBase) error
	Delete(ctx context.Context, key string) error
}
