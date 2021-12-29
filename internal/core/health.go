package core

import (
	"github.com/diazharizky/rest-otp-generator/pkg/health"
)

func (c *core) HealthCheck() health.HealthStatus {
	dbHealth := true
	if err := c.DB.Health(); err != nil {
		dbHealth = false
	}
	return health.HealthStatus{DB: dbHealth}
}
