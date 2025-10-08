//nolint:dupl // similar to other entity
package usecase

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/1nterdigital/aka-im-discover/internal/domain"
	model "github.com/1nterdigital/aka-im-discover/internal/model"
	discoveryCarousels "github.com/1nterdigital/aka-im-discover/internal/repository/discover/carousels"
	"github.com/1nterdigital/aka-im-tools/tracer"
)

type DiscoverCarouselsUseCase struct {
	discoverCarouselsRepo discoveryCarousels.Repository
}

func NewDiscoverCarouselsUseCase(discoverCarouselsRepo discoveryCarousels.Repository) *DiscoverCarouselsUseCase {
	return &DiscoverCarouselsUseCase{
		discoverCarouselsRepo: discoverCarouselsRepo,
	}
}

func (u *DiscoverCarouselsUseCase) Create(
	ctx context.Context,
	item *domain.DiscoverCarouselsAddReq,
) (resp *model.DiscoverCarousels, err error) {
	ctx, span := otel.Tracer(domain.TracerLevelUsecase).
		Start(ctx, tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	carousel := &model.DiscoverCarousels{
		Title:     item.Title,
		ImageURL:  item.ImageURL,
		LinkURL:   item.LinkURL,
		CreatedBy: item.CreatedBy,
		Position:  item.Position,
	}
	span.SetAttributes(
		attribute.String("title", carousel.Title),
		attribute.String("createdBy", carousel.CreatedBy),
	)

	err = u.discoverCarouselsRepo.Create(ctx, carousel)
	if err != nil {
		return nil, err
	}

	return carousel, nil
}

func (u *DiscoverCarouselsUseCase) Find(
	ctx context.Context, req *domain.DiscoverCarouselsFindReq,
) (resp []*model.DiscoverCarousels, count int64, err error) {
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
		attribute.Int64("carouselID", req.ID),
		attribute.String("title", req.Title),
		attribute.Int("page", int(req.Page)),
		attribute.Int("limit", int(req.Limit)),
	)

	resp, count, err = u.discoverCarouselsRepo.Find(ctx, req)
	return resp, count, err
}

func (u *DiscoverCarouselsUseCase) Delete(ctx context.Context, id int64, deletedBy string) (err error) {
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
		attribute.Int64("carouselID", id),
		attribute.String("deletedBy", deletedBy),
	)

	err = u.discoverCarouselsRepo.Delete(ctx, id, deletedBy)
	return err
}

func (u *DiscoverCarouselsUseCase) Edit(
	ctx context.Context, item *domain.DiscoverCarouselsEditReq,
) (resp *model.DiscoverCarousels, err error) {
	ctx, span := otel.Tracer(domain.TracerLevelUsecase).
		Start(ctx, tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	carousel := model.DiscoverCarousels{
		ID:        item.ID,
		Title:     item.Title,
		ImageURL:  item.ImageURL,
		LinkURL:   item.LinkURL,
		UpdatedBy: item.UpdatedBy,
		Position:  item.Position,
	}
	span.SetAttributes(
		attribute.Int64("carouselID", carousel.ID),
		attribute.String("title", carousel.Title),
		attribute.String("updatedBy", carousel.UpdatedBy),
	)

	resp, err = u.discoverCarouselsRepo.Edit(ctx, &carousel)
	return resp, err
}
