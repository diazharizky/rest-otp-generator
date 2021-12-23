package configs

import (
	"testing"
)

const (
	defaultHost = "0.0.0.0"
	defaultPort = 8080
)

func TestConfigDefaultValue(t *testing.T) {
	listenHost := Cfg.GetString("listen.host")
	if listenHost != defaultHost {
		t.Errorf("Listen host default value doesn't match")
	}

	listenPort := Cfg.GetInt("listen.port")
	if listenPort != defaultPort {
		t.Errorf("Listen port default value doesn't match")
	}
}
