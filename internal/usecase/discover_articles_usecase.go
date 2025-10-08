//nolint:dupl // similar to other entity
package usecase

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/1nterdigital/aka-im-discover/internal/domain"
	model "github.com/1nterdigital/aka-im-discover/internal/model"
	discoveryArticles "github.com/1nterdigital/aka-im-discover/internal/repository/discover/articles"
	"github.com/1nterdigital/aka-im-tools/tracer"
)

type DiscoverArticlesUseCase struct {
	discoverArticlesRepo discoveryArticles.Repository
}

func NewDiscoverArticlesUseCase(discoverArticlesRepo discoveryArticles.Repository) *DiscoverArticlesUseCase {
	return &DiscoverArticlesUseCase{
		discoverArticlesRepo: discoverArticlesRepo,
	}
}

func (u *DiscoverArticlesUseCase) Create(
	ctx context.Context, item *domain.DiscoverArticlesAddReq,
) (resp *model.DiscoverArticles, err error) {
	ctx, span := otel.Tracer(domain.TracerLevelUsecase).
		Start(ctx, tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	article := &model.DiscoverArticles{
		Title:     item.Title,
		ImageURL:  item.ImageURL,
		LinkURL:   item.LinkURL,
		CreatedBy: item.CreatedBy,
		Position:  item.Position,
	}
	span.SetAttributes(
		attribute.String("title", article.Title),
		attribute.String("createdBy", article.CreatedBy),
	)

	err = u.discoverArticlesRepo.Create(ctx, article)
	if err != nil {
		return nil, err
	}

	return article, nil
}
func (u *DiscoverArticlesUseCase) Find(
	ctx context.Context, req *domain.DiscoverArticlesFindReq,
) (resp []*model.DiscoverArticles, total int64, err error) {
	ctx, span := otel.Tracer(domain.TracerLevelUsecase).
		Start(ctx, tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	span.SetAttributes(
		attribute.Int64("articleID", req.ID),
		attribute.String("title", req.Title),
		attribute.Int("page", int(req.Page)),
		attribute.Int("limit", int(req.Limit)),
	)

	resp, total, err = u.discoverArticlesRepo.Find(ctx, req)
	return resp, total, err
}

func (u *DiscoverArticlesUseCase) Delete(ctx context.Context, id int64, deletedBy string) (err error) {
	ctx, span := otel.Tracer(domain.TracerLevelUsecase).
		Start(ctx, tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	span.SetAttributes(
		attribute.Int64("articleID", id),
		attribute.String("deletedBy", deletedBy),
	)

	err = u.discoverArticlesRepo.Delete(ctx, id, deletedBy)
	return err
}

func (u *DiscoverArticlesUseCase) Edit(
	ctx context.Context, item *domain.DiscoverArticlesEditReq,
) (resp *model.DiscoverArticles, err error) {
	ctx, span := otel.Tracer(domain.TracerLevelUsecase).
		Start(ctx, tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	article := model.DiscoverArticles{
		ID:        item.ID,
		Title:     item.Title,
		ImageURL:  item.ImageURL,
		LinkURL:   item.LinkURL,
		UpdatedBy: item.UpdatedBy,
		Position:  item.Position,
	}
	span.SetAttributes(
		attribute.Int64("articleID", item.ID),
		attribute.String("title", item.Title),
		attribute.String("updatedBy", item.UpdatedBy),
	)

	resp, err = u.discoverArticlesRepo.Edit(ctx, &article)
	return resp, err
}
