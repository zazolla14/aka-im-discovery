package repository

import (
	"gorm.io/gorm"

	"github.com/1nterdigital/aka-im-discover/internal/repository/discover/articles"
	"github.com/1nterdigital/aka-im-discover/internal/repository/discover/carousels"
	health "github.com/1nterdigital/aka-im-discover/internal/repository/health"
)

type Repository interface {
	Health() health.Repository
	DiscoverArticles() articles.Repository
	DiscoverCarousels() carousels.Repository
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Health() health.Repository {
	return health.New(r.db)
}

func (r *repository) DiscoverCarousels() carousels.Repository {
	return carousels.New(r.db)
}

func (r *repository) DiscoverArticles() articles.Repository {
	return articles.New(r.db)
}
