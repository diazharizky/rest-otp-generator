package interfaces

import (
	"os"
	"strconv"

	"github.com/diazharizky/rest-otp-generator/internal/application"
	"github.com/diazharizky/rest-otp-generator/internal/infrastructure/cache"
	"github.com/go-chi/chi"
)

const (
	CacheDefaultHost = "0.0.0.0"
	CacheDefaultPort = "6379"
	CacheDefaultPass = ""
	CacheDefaultDB   = "0"

	EnvCacheHost = "CACHE_HOST"
	EnvCachePort = "CACHE_PORT"
	EnvCachePass = "CACHE_PASSWORD"
	EnvCacheDB   = "CACHE_DB"
)

func Router() (r chi.Router) {
	c := cache.NewCache(CacheConfig())
	ha := application.NewHealthApp(&c.HealthCache)
	oa := application.NewOTPApp(&c.OTPCache)

	hh := NewHealthHandler(&ha)
	oh := NewOTPHandler(&oa)

	r = chi.NewRouter()
	r.Mount("/health", hh.getHandler())
	r.Mount("/otp", oh.getHandler())
	return
}

func CacheConfig() (string, string, string, int) {
	host := os.Getenv(EnvCacheHost)
	if len(host) <= 0 {
		host = CacheDefaultHost
	}

	port := os.Getenv(EnvCachePort)
	if len(port) <= 0 {
		port = CacheDefaultPort
	}

	passwd := os.Getenv(EnvCachePass)
	dbString := os.Getenv(EnvCacheDB)
	if len(dbString) <= 0 {
		dbString = CacheDefaultDB
	}
	db, err := strconv.Atoi(dbString)
	if err != nil {
		panic(err)
	}

	return host, port, passwd, db
}
