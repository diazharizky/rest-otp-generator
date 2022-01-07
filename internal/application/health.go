package application

import (
	"github.com/diazharizky/rest-otp-generator/internal/domain"
	"github.com/diazharizky/rest-otp-generator/internal/domain/repository"
)

type HealthAppInterface interface {
	HealthCheck() domain.Health
}

type healthApp struct {
	r repository.HealthRepository
}

func NewHealthApp(hr repository.HealthRepository) healthApp {
	return healthApp{hr}
}

func (h *healthApp) HealthCheck() domain.Health {
	cacheHealth := true
	if err := h.r.CacheHealth(); err != nil {
		cacheHealth = false
	}
	return domain.Health{Cache: cacheHealth}
}
