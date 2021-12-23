package db

import (
	"context"
	"time"
)

type Database interface {
	Health() error
	Get(context.Context, string) ([]byte, error)
	Set(context.Context, string, interface{}, time.Duration) error
	Delete(ctx context.Context, key string) error
}
