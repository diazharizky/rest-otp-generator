package health

import (
	"github.com/diazharizky/rest-otp-generator/internal/types"
)

type HealthStatus struct {
	Redis string `json:"redis"`
}

type core struct {
	Redis types.DBService
}
