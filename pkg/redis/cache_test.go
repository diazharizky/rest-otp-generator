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
	"github.com/diazharizky/rest-otp-generator/internal/db"
	"github.com/diazharizky/rest-otp-generator/pkg/otp"
	cache "github.com/diazharizky/rest-otp-generator/pkg/redis"
	"github.com/stretchr/testify/suite"
)

const (
	defaultPeriod      = 60
	defaultDigits      = 4
	defaultMaxAttemtps = 4
)

var otpV otp.OTPBase

func init() {
	otpV = otp.OTPBase{
		Period:      defaultPeriod,
		Digits:      defaultDigits,
		MaxAttempts: defaultMaxAttemtps,
	}
}

type HandlerSuite struct {
	RedisSuite
}

func TestRedisSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for Redis repository")
	}

	cfg := cache.Config{Config: db.Config{
		Host:     configs.Cfg.GetString("cache.host"),
		Port:     configs.Cfg.GetString("cache.port"),
		Password: configs.Cfg.GetString("cache.password"),
		DB:       configs.Cfg.GetInt("cache.db"),
	}}
	HandlerSuiteTest := &HandlerSuite{RedisSuite{Config: cfg}}
	suite.Run(t, HandlerSuiteTest)
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
	return client.Set(ctx, key, jbt, defaultPeriod*time.Second).Err()
}

func (r *HandlerSuite) TestHealth() {
	err := cache.Handler.Health()
	assert.Nil(r.T(), err)
}

func (r *HandlerSuite) TestGet() {
	testKey := "3f6cf12b-2a16-4a0a-97d7-2c5c2152c7db"
	err := seedDB(r.Client, testKey, otpV)
	require.NoError(r.T(), err)

	ctx := context.Background()
	jbt, err := cache.Handler.Get(ctx, testKey)
	require.NoError(r.T(), err)
	require.NotNil(r.T(), jbt)

	var storedData otp.OTPBase
	err = json.Unmarshal(jbt, &storedData)
	require.NoError(r.T(), err)

	assert.Equal(r.T(), storedData.Period, otpV.Period)
	assert.Equal(r.T(), storedData.Digits, otpV.Digits)
	assert.Equal(r.T(), storedData.MaxAttempts, otpV.MaxAttempts)
}

func (r *HandlerSuite) TestSet() {
	ctx := context.Background()
	testKey := "6cd76df0-f151-4f06-b17e-8235508d0273"
	err := cache.Handler.Set(ctx, testKey, otpV, time.Duration(otpV.Period)*time.Second)
	require.NoError(r.T(), err)

	jbt, err := getItemByKey(r.Client, testKey)
	require.NoError(r.T(), err)
	require.NotNil(r.T(), jbt)

	var storedData otp.OTPBase
	err = json.Unmarshal(jbt, &storedData)
	require.NoError(r.T(), err)

	assert.Equal(r.T(), storedData.Period, otpV.Period)
	assert.Equal(r.T(), storedData.Digits, otpV.Digits)
	assert.Equal(r.T(), storedData.MaxAttempts, otpV.MaxAttempts)
}

func (r *HandlerSuite) TestDelete() {
	testKey := "8e1e8d58-6c0f-4b8c-bc7e-1dc3442b327f"
	err := seedDB(r.Client, testKey, otpV)
	require.NoError(r.T(), err)

	ctx := context.Background()
	jbt, err := cache.Handler.Get(ctx, testKey)
	require.NoError(r.T(), err)
	require.NotNil(r.T(), jbt)

	var storedData otp.OTPBase
	err = json.Unmarshal(jbt, &storedData)
	require.NoError(r.T(), err)

	assert.Equal(r.T(), storedData.Period, otpV.Period)
	assert.Equal(r.T(), storedData.Digits, otpV.Digits)
	assert.Equal(r.T(), storedData.MaxAttempts, otpV.MaxAttempts)

	err = cache.Handler.Delete(ctx, testKey)
	require.NoError(r.T(), err)

	jbt, err = cache.Handler.Get(ctx, testKey)
	require.NoError(r.T(), err)
	require.Nil(r.T(), jbt)
}
