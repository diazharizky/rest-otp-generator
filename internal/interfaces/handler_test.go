package interfaces_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/diazharizky/rest-otp-generator/internal/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestCacheConfig(t *testing.T) {
	dbString := os.Getenv(interfaces.EnvCacheDB)
	if len(dbString) <= 0 {
		dbString = interfaces.EnvCacheDB
	}
	dbInt, _ := strconv.Atoi(dbString)

	host, port, pass, db := interfaces.CacheConfig()
	hostOneOf := host == interfaces.CacheDefaultHost || host == os.Getenv(interfaces.EnvCacheHost)
	portOneOf := port == interfaces.CacheDefaultPort || port == os.Getenv(interfaces.EnvCachePort)
	passOneOf := pass == interfaces.CacheDefaultPass || pass == os.Getenv(interfaces.EnvCachePort)
	dbOneOf := db == dbInt

	assert.True(t, hostOneOf)
	assert.True(t, portOneOf)
	assert.True(t, passOneOf)
	assert.True(t, dbOneOf)
}
