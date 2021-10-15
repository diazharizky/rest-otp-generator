package types

import (
	"context"

	"github.com/diazharizky/rest-otp-generator/pkg/otp"
)

type Service interface {
	Health() error
}

type DBService interface {
	Service

	Get(context.Context, otp.OTP) error
	Upsert(context.Context, otp.OTP) error
	Delete(ctx context.Context, id string) error
}
