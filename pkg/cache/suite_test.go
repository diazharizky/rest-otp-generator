package cache_test

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/suite"
)

type CacheSuite struct {
	suite.Suite
	Host     string
	Port     string
	Password string
	DB       int
	Client   *redis.Client
}

func (r *CacheSuite) SetupSuite() {
	addr := fmt.Sprintf("%s:%s", r.Host, r.Port)
	r.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: r.Password,
		DB:       r.DB,
	})
}

func (r *CacheSuite) TearDownSuite() {
	r.Client.Close()
}
