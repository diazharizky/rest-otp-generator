package redis_test

import (
	"fmt"

	myRedis "github.com/diazharizky/rest-otp-generator/pkg/redis"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/suite"
)

type RedisSuite struct {
	suite.Suite
	myRedis.RedisConfig
}

func (r *RedisSuite) SetupSuite() {
	addr := fmt.Sprintf("%s:%s", r.Host, r.Port)
	r.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: r.Password,
		DB:       r.DB,
	})
}

func (r *RedisSuite) TearDownSuite() {
	r.Client.Close()
}
