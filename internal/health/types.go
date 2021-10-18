package health

import (
	"github.com/diazharizky/rest-otp-generator/internal/types"
)

type core struct {
	DB types.DBService
}

type healthStatus struct {
	DB string `json:"db"`
}
