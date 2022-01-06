package repository

type HealthRepository interface {
	CacheHealth() error
}
