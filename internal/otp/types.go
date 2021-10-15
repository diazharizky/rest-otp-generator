package otp

import (
	"context"

	"github.com/diazharizky/rest-otp-generator/internal/types"
	"github.com/diazharizky/rest-otp-generator/pkg/otp"
)

type DBService interface {
	types.Service

	Get(context.Context, otp.OTP) error
	Upsert(context.Context, otp.OTP) error
	Delete(ctx context.Context, id string) error
}

type core struct {
	DB DBService
}
