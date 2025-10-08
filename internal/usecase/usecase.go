package usecase

import (
	"github.com/1nterdigital/aka-im-discover/internal/repository"
)

type UseCase struct {
	Health            *HealthUseCase
	DiscoverArticles  *DiscoverArticlesUseCase
	DiscoverCarousels *DiscoverCarouselsUseCase
}

func New(repo repository.Repository) (*UseCase, error) {
	healthUsecase := NewHealthUseCase(
		repo.Health(),
	)

	discoverArticlesUsecase := NewDiscoverArticlesUseCase(
		repo.DiscoverArticles(),
	)

	discoverCarouselsUsecase := NewDiscoverCarouselsUseCase(
		repo.DiscoverCarousels(),
	)

	return &UseCase{
		Health:            healthUsecase,
		DiscoverArticles:  discoverArticlesUsecase,
		DiscoverCarousels: discoverCarouselsUsecase,
	}, nil
}
