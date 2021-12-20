package configs

import (
	"testing"
)

func TestConfigDefaultValue(t *testing.T) {
	listenHost := Cfg.GetString("listen.host")
	if listenHost != "0.0.0.0" {
		t.Errorf("Listen host default value doesn't match")
	}

	listenPort := Cfg.GetString("listen.port")
	if listenPort != "8080" {
		t.Errorf("Listen port default value doesn't match")
	}
}
