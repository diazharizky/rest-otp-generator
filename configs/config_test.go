package configs_test

import (
	"testing"

	"github.com/diazharizky/rest-otp-generator/configs"

	"github.com/stretchr/testify/assert"
)

const (
	defaultHost = "0.0.0.0"
	defaultPort = 8080

	defaultCacheHost     = "0.0.0.0"
	defaultCachePort     = 6379
	defaultCachePassword = ""
	defaultCacheDB       = 0
)

func TestConfigDefaultValue(t *testing.T) {
	listenHost := configs.Cfg.GetString("listen.host")
	assert.Equal(t, listenHost, defaultHost, "Listen host default value doesn't match")

	listenPort := configs.Cfg.GetInt("listen.port")
	assert.Equal(t, listenPort, defaultPort, "Listen port default value doesn't match")

	cacheHost := configs.Cfg.GetString("cache.host")
	assert.Equal(t, cacheHost, defaultCacheHost, "Cache's host default value doesn't match")

	cachePort := configs.Cfg.GetInt("cache.port")
	assert.Equal(t, cachePort, defaultCachePort, "Cache's port default value doesn't match")

	cachePassword := configs.Cfg.GetString("cache.password")
	assert.Equal(t, cachePassword, defaultCachePassword, "Cache's password default value doesn't match")

	cacheDB := configs.Cfg.GetInt("cache.db")
	assert.Equal(t, cacheDB, defaultCacheDB, "Cache's DB default value doesn't match")
}
