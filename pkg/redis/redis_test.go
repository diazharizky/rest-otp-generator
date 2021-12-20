package redis

import (
	"testing"
)

var svc Service

func init() {
	svc = Service{Client: Connect(GetCfg())}
}

func TestGetConnection(t *testing.T) {
	if err := svc.Health(); err != nil {
		t.Errorf("DB connection failures to establish.")
	}
}
