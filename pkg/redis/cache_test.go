package redis_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	myRedis "github.com/diazharizky/rest-otp-generator/pkg/redis"
	"github.com/stretchr/testify/suite"
)

type redisHandlerSuite struct {
	RedisSuite // Embed RedisSuite struct
}

func TestRedisSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for redis repository")
	}
	cfg := myRedis.RedisConfig{
		Host:     configs.Cfg.GetString("redis.host"),
		Port:     configs.Cfg.GetString("redis.port"),
		Password: configs.Cfg.GetString("redis.password"),
		DB:       configs.Cfg.GetInt("redis.db"),
	}
	redisHandlerSuiteTest := &redisHandlerSuite{RedisSuite{RedisConfig: cfg}}
	suite.Run(t, redisHandlerSuiteTest)
}

func getItemByKey(client *redis.Client, key string) ([]byte, error) {
	ctx := context.Background()
	return client.Get(ctx, key).Bytes()
}

func seedDB(client *redis.Client, key string, value interface{}) error {
	jbt, err := json.Marshal(value)
	if err != nil {
		return err
	}
	ctx := context.Background()
	return client.Set(ctx, key, jbt, time.Second*30).Err()
}

func (r *redisHandlerSuite) TestGet() {
	testKey := "3f6cf12b-2a16-4a0a-97d7-2c5c2152c7db"
	otpV := otp.OTPBase{
		Key:         testKey,
		Period:      60 * time.Second,
		Digits:      4,
		MaxAttempts: 4,
	}
	err := seedDB(r.Client, testKey, otpV)
	require.NoError(r.T(), err)

	repo := myRedis.GetHandler(r.Client)
	ctx := context.Background()
	jbt, err := repo.Get(ctx, testKey)
	require.NoError(r.T(), err)

	var res otp.OTPBase
	err = json.Unmarshal(jbt, &res)
	require.NoError(r.T(), err)

	assert.Equal(r.T(), otpV.Key, res.Key)
	assert.Equal(r.T(), otpV.Period, res.Period)
	assert.Equal(r.T(), otpV.Digits, res.Digits)
}

func (r *redisHandlerSuite) TestSet() {
	testKey := "6cd76df0-f151-4f06-b17e-8235508d0273"
	repo := myRedis.GetHandler(r.Client)
	otpV := otp.OTPBase{
		Key:         testKey,
		Period:      60 * time.Second,
		Digits:      4,
		MaxAttempts: 4,
	}
	ctx := context.Background()
	err := repo.Set(ctx, testKey, otpV, 300*time.Second)
	require.NoError(r.T(), err)

	jbt, err := getItemByKey(r.Client, testKey)
	require.NoError(r.T(), err)
	require.NotNil(r.T(), jbt)

	var storedData otp.OTPBase
	err = json.Unmarshal(jbt, &storedData)
	require.NoError(r.T(), err)

	assert.Equal(r.T(), storedData.Key, otpV.Key)
	assert.Equal(r.T(), storedData.Period, otpV.Period)
	assert.Equal(r.T(), storedData.Digits, otpV.Digits)
	assert.Equal(r.T(), storedData.MaxAttempts, otpV.MaxAttempts)
}
