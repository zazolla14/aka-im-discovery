package usecase

import (
	"context"
	"log"

	"github.com/1nterdigital/aka-im-discover/internal/repository/health"
)

type HealthUseCase struct {
	healthRepo health.Repository
}

func NewHealthUseCase(healthRepo health.Repository) *HealthUseCase {
	return &HealthUseCase{
		healthRepo: healthRepo,
	}
}

func (u *HealthUseCase) HealthCheck(ctx context.Context) error {
	err := u.healthRepo.HealthCheck(ctx)
	if err != nil {
		log.Println("Health check failed:", err)
		return err
	}

	return nil
}
