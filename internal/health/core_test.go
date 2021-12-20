package health

import (
	"testing"
)

func TestHealthCheck(t *testing.T) {
	h := c.healthCheck()
	if h.Db == false {
		t.Errorf("DB is not ready.")
	}
}
