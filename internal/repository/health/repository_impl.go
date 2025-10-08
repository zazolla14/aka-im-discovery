package health

import (
	"context"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

//nolint:revive // HealthCheck implements health.Repository
func (r *repositoryImpl) HealthCheck(_ context.Context) error {
	return nil
}
