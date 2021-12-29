package core

import (
	"github.com/diazharizky/rest-otp-generator/pkg/db"
	cache "github.com/diazharizky/rest-otp-generator/pkg/redis"
)

type core struct {
	DB db.Database
}

var Core core

func init() {
	Core.DB = &cache.Handler
}
