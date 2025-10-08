package articles

import (
	"context"

	"github.com/1nterdigital/aka-im-discover/internal/domain"
	model "github.com/1nterdigital/aka-im-discover/internal/model"
)

type Repository interface {
	Create(ctx context.Context, item *model.DiscoverArticles) (err error)
	Find(
		ctx context.Context, req *domain.DiscoverArticlesFindReq,
	) (resp []*model.DiscoverArticles, count int64, err error)
	Delete(ctx context.Context, id int64, deletedBy string) (err error)
	Edit(
		ctx context.Context, article *model.DiscoverArticles,
	) (resp *model.DiscoverArticles, err error)
}
