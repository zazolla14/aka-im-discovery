package carousels

import (
	"context"

	"github.com/1nterdigital/aka-im-discover/internal/domain"
	model "github.com/1nterdigital/aka-im-discover/internal/model"
)

type Repository interface {
	Create(ctx context.Context, item *model.DiscoverCarousels) error
	Find(
		ctx context.Context, req *domain.DiscoverCarouselsFindReq,
	) (resp []*model.DiscoverCarousels, count int64, err error)
	Delete(ctx context.Context, id int64, deletedBy string) error
	Edit(
		ctx context.Context, carousel *model.DiscoverCarousels,
	) (resp *model.DiscoverCarousels, err error)
}
