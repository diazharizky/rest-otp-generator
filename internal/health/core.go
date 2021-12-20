package health

import (
	"github.com/diazharizky/rest-otp-generator/internal/db"
	"github.com/diazharizky/rest-otp-generator/pkg/redis"
)

type core struct {
	Db db.Database
}

type healthStatus struct {
	Db bool `json:"database"`
}

var c core

func init() {
	client := redis.Connect(redis.GetCfg())
	c.Db = &redis.Service{Client: client}
}

func (c *core) healthCheck() healthStatus {
	dbHealth := true
	if err := c.Db.Health(); err != nil {
		dbHealth = false
	}
	return healthStatus{Db: dbHealth}
}
