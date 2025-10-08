package http

import (
	"github.com/1nterdigital/aka-im-discover/internal/service"
	"github.com/1nterdigital/aka-im-discover/internal/usecase"
)

type DiscoverHandler struct {
	healthUsecase            *usecase.HealthUseCase
	discoverArticlesUsecase  *usecase.DiscoverArticlesUseCase
	discoverCarouselsUsecase *usecase.DiscoverCarouselsUseCase
}

func NewDiscoverHandler(u *service.Api) *DiscoverHandler {
	return &DiscoverHandler{
		healthUsecase:            u.HealthUseCase().Health,
		discoverArticlesUsecase:  u.DiscoverUseCase().DiscoverArticles,
		discoverCarouselsUsecase: u.DiscoverUseCase().DiscoverCarousels,
	}
}
