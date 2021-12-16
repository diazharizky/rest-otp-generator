package redis

import (
	"os"
	"strconv"
	"testing"
)

var svc Service

func init() {
	dbIndex, _ := strconv.ParseInt(os.Getenv("OTPGEN_REDIS_DB"), 10, 0)
	cfg := Cfg{
		Host:     os.Getenv("OTPGEN_REDIS_HOST"),
		Port:     os.Getenv("OTPGEN_REDIS_PORT"),
		Database: int(dbIndex),
	}
	svc = Service{Client: Connect(cfg)}
}

func TestGetConnection(t *testing.T) {
	if err := svc.Health(); err != nil {
		t.Errorf("DB connection failures to establish.")
	}
}
