package carousels

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"gorm.io/gorm"

	"github.com/1nterdigital/aka-im-discover/internal/domain"
	model "github.com/1nterdigital/aka-im-discover/internal/model"
	"github.com/1nterdigital/aka-im-tools/tracer"
)

type repositoryImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) Create(ctx context.Context, carousel *model.DiscoverCarousels) (err error) {
	ctx, span := otel.Tracer(domain.TracerLevelRepository).
		Start(ctx, tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	span.SetAttributes(
		attribute.String("title", carousel.Title),
		attribute.String("createdBy", carousel.CreatedBy),
	)

	var item model.DiscoverCarousels
	err = r.db.WithContext(ctx).
		Where("title = ? AND is_active = ? AND deleted_at IS NULL", carousel.Title, true).
		First(&item).Error

	if err == nil {
		return fmt.Errorf("title already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	err = r.db.WithContext(ctx).Create(&carousel).Error
	return err
}

func (r *repositoryImpl) Find(
	ctx context.Context, req *domain.DiscoverCarouselsFindReq,
) (resp []*model.DiscoverCarousels, count int64, err error) {
	ctx, span := otel.Tracer(domain.TracerLevelRepository).
		Start(ctx, tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	var (
		items  []*model.DiscoverCarousels
		total  int64
		offset = (req.Page - 1) * req.Limit
	)

	span.SetAttributes(
		attribute.Int64("id", req.ID),
		attribute.String("title", req.Title),
		attribute.Int("page", int(req.Page)),
		attribute.Int("limit", int(req.Limit)),
	)

	if req.SortBy == "position" {
		req.SortBy = "position IS NULL, position"
	}

	query := r.db.WithContext(ctx).
		Model(&model.DiscoverCarousels{}).
		Where("is_active = ? AND deleted_at IS NULL", true).
		Order(req.SortBy + " " + req.Order)

	if req.ID != 0 {
		query = query.Where("id = ?", req.ID)
	}

	if req.Title != "" {
		query = query.Where("title LIKE ?", "%"+req.Title+"%")
	}

	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	err = query.Limit(int(req.Limit)).
		Offset(int(offset)).
		Find(&items).Error

	return items, total, err
}

func (r *repositoryImpl) Delete(ctx context.Context, id int64, deletedBy string) (err error) {
	ctx, span := otel.Tracer(domain.TracerLevelRepository).
		Start(ctx, tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	span.SetAttributes(
		attribute.Int64("id", id),
		attribute.String("deletedBy", deletedBy),
	)

	var item model.DiscoverCarousels
	err = r.db.WithContext(ctx).
		Where("id = ? AND is_active = ? AND deleted_at IS NULL", id, true).
		First(&item).Error
	if err != nil {
		return err
	}

	err = r.db.WithContext(ctx).
		Model(&item).
		Updates(map[string]interface{}{
			"deleted_by": deletedBy,
			"is_active":  false,
			"deleted_at": time.Now(),
		}).Error
	return err
}

func (r *repositoryImpl) Edit(
	ctx context.Context, carousel *model.DiscoverCarousels,
) (resp *model.DiscoverCarousels, err error) {
	ctx, span := otel.Tracer(domain.TracerLevelRepository).
		Start(ctx, tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	span.SetAttributes(
		attribute.Int64("id", carousel.ID),
		attribute.String("title", carousel.Title),
		attribute.String("updatedBy", carousel.UpdatedBy),
	)

	var item model.DiscoverCarousels
	query := r.db.WithContext(ctx).
		Where("is_active = ? AND deleted_at IS NULL", true)

	err = query.Where("id = ?", carousel.ID).First(&item).Error
	if err != nil {
		return nil, err
	}

	var dup model.DiscoverCarousels
	err = query.Where("title = ?", carousel.Title).First(&dup).Error
	if err == nil {
		return nil, fmt.Errorf("title already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	updates := map[string]interface{}{}

	if carousel.Title != "" {
		updates["title"] = carousel.Title
	}
	if carousel.ImageURL != "" {
		updates["image_url"] = carousel.ImageURL
	}
	if carousel.LinkURL != "" {
		updates["link_url"] = carousel.LinkURL
	}

	if carousel.Position != nil {
		updates["position"] = carousel.Position
	}

	updates["updated_by"] = carousel.UpdatedBy

	err = r.db.WithContext(ctx).Model(&item).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}
