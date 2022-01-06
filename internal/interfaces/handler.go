package interfaces

import (
	"os"
	"strconv"

	"github.com/diazharizky/rest-otp-generator/internal/application"
	"github.com/diazharizky/rest-otp-generator/internal/infrastructure/cache"
	"github.com/go-chi/chi"
)

func Router() (r chi.Router) {
	c := cache.NewCache(cacheConfig())
	ha := application.NewHealthApp(&c.HealthCache)
	oa := application.NewOTPApp(&c.OTPCache)

	hh := newHealthHandler(&ha)
	oh := newOTPHandler(&oa)

	r = chi.NewRouter()
	r.Mount("/health", hh.getHandler())
	r.Mount("/otp", oh.getHandler())
	return
}

func cacheConfig() (string, string, string, int) {
	host := os.Getenv("CACHE_HOST")
	if len(host) <= 0 {
		host = "0.0.0.0"
	}

	port := os.Getenv("CACHE_PORT")
	if len(port) <= 0 {
		port = "6379"
	}

	passwd := os.Getenv("CACHE_PASSWORD")
	dbString := os.Getenv("CACHE_DB")
	if len(dbString) <= 0 {
		dbString = "0"
	}
	db, err := strconv.Atoi(dbString)
	if err != nil {
		panic(err)
	}

	return host, port, passwd, db
}
