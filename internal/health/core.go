package health

import (
	"fmt"

	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/diazharizky/rest-otp-generator/internal/db"
	"github.com/diazharizky/rest-otp-generator/pkg/redis"

	myRedis "github.com/go-redis/redis/v8"
)

type core struct {
	DB db.Database
}

type healthStatus struct {
	DB bool `json:"database"`
}

var c core

func init() {
	dbHost := configs.Cfg.GetString("redis.host")
	dbPort := configs.Cfg.GetString("redis.port")
	addr := fmt.Sprintf("%s:%s", dbHost, dbPort)
	client := myRedis.NewClient(&myRedis.Options{
		Addr:     addr,
		Password: configs.Cfg.GetString("redis.password"),
		DB:       configs.Cfg.GetInt("redis.db"),
	})
	c.DB = redis.GetHandler(client)
}

func (c *core) healthCheck() healthStatus {
	dbHealth := true
	if err := c.DB.Health(); err != nil {
		dbHealth = false
	}
	return healthStatus{DB: dbHealth}
}
