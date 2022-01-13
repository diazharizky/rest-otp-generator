package cache_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/diazharizky/rest-otp-generator/pkg/cache"
	"github.com/go-redis/redis/v8"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	testKey       = "some-test-key"
	cacheDuration = 30 * time.Second
)

type book struct {
	Title     string
	Author    string
	Publisher string
}

type cacheHandlerSuite struct {
	CacheSuite
}

var b book

func init() {
	b = book{
		Title:     "Jungle Book",
		Author:    "Fulan",
		Publisher: "Gramedia",
	}
}

func TestRedisSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for redis repository")
	}

	host, port, passwd, db := cache.CacheConfig()
	cacheHandlerSuiteTest := &cacheHandlerSuite{
		CacheSuite{
			Host:     host,
			Port:     port,
			Password: passwd,
			DB:       db,
		},
	}
	suite.Run(t, cacheHandlerSuiteTest)
}

func getItemByKey(client *redis.Client, key string) ([]byte, error) {
	ctx := context.Background()
	return client.Get(ctx, key).Bytes()
}

func seedItem(client *redis.Client, key string, value interface{}) error {
	byt, err := json.Marshal(value)
	if err != nil {
		return err
	}
	ctx := context.Background()
	return client.Set(ctx, key, byt, cacheDuration).Err()
}

func (r *cacheHandlerSuite) TestSet() {
	err := cache.Set(r.Client, testKey, b, cacheDuration)
	require.NoError(r.T(), err)

	byt, err := getItemByKey(r.Client, testKey)
	require.NoError(r.T(), err)
	require.NotNil(r.T(), byt)
	var storedData book
	err = json.Unmarshal(byt, &storedData)
	require.NoError(r.T(), err)

	assert.Equal(r.T(), b.Title, storedData.Title)
	assert.Equal(r.T(), b.Author, storedData.Author)
	assert.Equal(r.T(), b.Publisher, storedData.Publisher)
}
func (r *cacheHandlerSuite) TestGet() {
	err := seedItem(r.Client, testKey, b)
	require.NoError(r.T(), err)

	byt, err := cache.Get(r.Client, testKey)
	require.NoError(r.T(), err)
	var storedData book
	err = json.Unmarshal(byt, &storedData)
	require.NoError(r.T(), err)

	assert.Equal(r.T(), b.Title, storedData.Title)
	assert.Equal(r.T(), b.Author, storedData.Author)
	assert.Equal(r.T(), b.Publisher, storedData.Publisher)
}
