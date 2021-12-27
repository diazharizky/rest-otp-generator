package health

import (
	"github.com/diazharizky/rest-otp-generator/internal/db"
	cache "github.com/diazharizky/rest-otp-generator/pkg/redis"
)

type core struct {
	DB db.Database
}

type healthStatus struct {
	DB bool `json:"database"`
}

var c core

func init() {
	c.DB = &cache.Handler
}

func (c *core) healthCheck() healthStatus {
	dbHealth := true
	if err := c.DB.Health(); err != nil {
		dbHealth = false
	}
	return healthStatus{DB: dbHealth}
}
