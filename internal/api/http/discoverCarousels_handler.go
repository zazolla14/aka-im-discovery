package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/1nterdigital/aka-im-discover/internal/domain"
	entity "github.com/1nterdigital/aka-im-discover/internal/model"
	"github.com/1nterdigital/aka-im-tools/apiresp"
	"github.com/1nterdigital/aka-im-tools/errs"
	"github.com/1nterdigital/aka-im-tools/log"
	"github.com/1nterdigital/aka-im-tools/mcontext"
	"github.com/1nterdigital/aka-im-tools/tracer"
)

// CreateCarousel create a carousel
//
// @Summary Create a new carousel
// @Description Creates a new carousel with the given JSON payload
// @Tags DiscoverCarousels
// @Accept json
// @Produce json
// @Param request body domain.DiscoverCarouselsAddReq true "Carousel request"
// @Success 200 {object} domain.DiscoverCarousels "Created carousel"
// @Failure 400 {object} apiresp.ApiResponse "Invalid json payload bad request"
// @Failure 500 {object} apiresp.ApiResponse "Internal server error"
// @Router /bo/discover/carousels/add [post]
// @Security ApiKeyAuth
func (h *DiscoverHandler) CreateCarousel(c *gin.Context) {
	var (
		req domain.DiscoverCarouselsAddReq
		err error
	)

	ctx, span := otel.Tracer(domain.TracerLevelHandler).
		Start(c.Request.Context(), tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.ZError(ctx, "an error occurred while CreateCarousel", err)
		}
		span.End()
	}()

	span.SetAttributes(
		attribute.String("userID", mcontext.GetOpUserID(c)),
		attribute.String("platformID", mcontext.GetOpUserPlatform(c)),
		attribute.String("operationID", mcontext.GetOperationID(c)),
	)

	req.CreatedBy, err = getOperatedByUser(c, req.CreatedBy)
	if err != nil {
		apiresp.GinError(c, err)
		return
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		err = errs.ErrArgs.WrapMsg("invalid json payload " + http.StatusText(http.StatusBadRequest))
		apiresp.GinError(c, err)
		return
	}
	carousel, err := h.discoverCarouselsUsecase.Create(ctx, &req)
	if err != nil {
		apiresp.GinError(c, err)
		return
	}
	apiresp.GinSuccess(c, carousel)
}

// FindCarousels Get paginated list of carousels
//
// @Summary Get paginated list of carousels
// @Description Retrieves a paginated list of carousels
// @Tags DiscoverCarousels
// @Accept json
// @Produce json
// @Param id query int false "carousel id" default("0")
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(10)
// @Param title query string false "Title name" default("")
// @Param sortBy query string false "Sort by" default("position")
// @Param order query string false "Order by" default("ASC")
// @Success 200 {array} domain.DiscoverCarousels "List of carousels"
// @Failure 400 {object} apiresp.ApiResponse "Invalid pagination parameters"
// @Failure 500 {object} apiresp.ApiResponse "Internal server error"
// @Router /bo/discover/carousels/find [get]
// @Security ApiKeyAuth
func (h *DiscoverHandler) FindCarousels(c *gin.Context) {
	var err error
	ctx, span := otel.Tracer(domain.TracerLevelHandler).
		Start(c.Request.Context(), tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.ZError(ctx, "an error occurred while FindCarousels", err)
		}
		span.End()
	}()

	span.SetAttributes(
		attribute.String("userID", mcontext.GetOpUserID(c)),
		attribute.String("platformID", mcontext.GetOpUserPlatform(c)),
		attribute.String("operationID", mcontext.GetOperationID(c)),
	)

	titleStr := c.DefaultQuery("title", "")

	id, err := parseIDParam(c.DefaultQuery("id", "0"))
	if err != nil {
		apiresp.GinError(c, err)
		return
	}

	page, limit, err := parsePaginationParams(c)
	if err != nil {
		apiresp.GinError(c, err)
		return
	}

	sortByStr, orderByStr, err := parseSortOrderParams(c)
	if err != nil {
		apiresp.GinError(c, err)
		return
	}

	req := domain.DiscoverCarouselsFindReq{
		ID:     id,
		Page:   page,
		Limit:  limit,
		Title:  titleStr,
		SortBy: sortByStr,
		Order:  orderByStr,
	}

	carousels, total, err := h.discoverCarouselsUsecase.Find(ctx, &req)
	if err != nil {
		apiresp.GinError(c, err)
		return
	}

	apiresp.GinSuccess(c, gin.H{
		"total": total,
		"data":  carousels,
	})
}

// DeleteCarousel Delete a carousel
//
// @Summary Delete a carousel
// @Description Deletes a carousel by its ID
// @Tags DiscoverCarousels
// @Accept json
// @Produce json
// @Param request body domain.DiscoverCarouselsDeleteReq true "Delete request"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} apiresp.ApiResponse "Invalid json payload bad request"
// @Failure 500 {object} apiresp.ApiResponse "Internal server error"
// @Router /discover/carousels/delete [delete]
// @Security ApiKeyAuth
func (h *DiscoverHandler) DeleteCarousel(c *gin.Context) {
	var (
		req domain.DiscoverCarouselsDeleteReq
		err error
	)

	ctx, span := otel.Tracer(domain.TracerLevelHandler).
		Start(c.Request.Context(), tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.ZError(ctx, "an error occurred while DeleteCarousel", err)
		}
		span.End()
	}()

	span.SetAttributes(
		attribute.String("userID", mcontext.GetOpUserID(c)),
		attribute.String("platformID", mcontext.GetOpUserPlatform(c)),
		attribute.String("operationID", mcontext.GetOperationID(c)),
	)

	req.DeletedBy, err = getOperatedByUser(c, req.DeletedBy)
	if err != nil {
		apiresp.GinError(c, err)
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		err = errs.ErrArgs.WrapMsg("invalid json payload " + http.StatusText(http.StatusBadRequest))
		apiresp.GinError(c, err)
		return
	}

	err = h.discoverCarouselsUsecase.Delete(ctx, req.ID, req.DeletedBy)
	if err != nil {
		apiresp.GinError(c, err)
		return
	}

	apiresp.GinSuccess(c, "deleted")
}

// EditCarousel Edit a carousel
//
// @Summary Edit a carousel
// @Description Updates an existing carousel
// @Tags DiscoverCarousels
// @Accept json
// @Produce json
// @Param request body domain.DiscoverCarouselsEditReq true "Edit request"
// @Success 200 {string} string "updated article"
// @Failure 400 {object} apiresp.ApiResponse "Invalid request payload"
// @Failure 500 {object} apiresp.ApiResponse "Internal server error"
// @Router /discover/carousels/edit [post]
// @Security ApiKeyAuth
func (h *DiscoverHandler) EditCarousel(c *gin.Context) {
	var (
		req domain.DiscoverCarouselsEditReq
		err error
	)

	ctx, span := otel.Tracer(domain.TracerLevelHandler).
		Start(c.Request.Context(), tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.ZError(ctx, "an error occurred while EditCarousel", err)
		}
		span.End()
	}()

	span.SetAttributes(
		attribute.String("userID", mcontext.GetOpUserID(c)),
		attribute.String("platformID", mcontext.GetOpUserPlatform(c)),
		attribute.String("operationID", mcontext.GetOperationID(c)),
	)

	if req.UpdatedBy == "" {
		userID := mcontext.GetOpUserID(c)
		if userID == "" {
			apiresp.GinError(c, errs.ErrUserIDEmpty.WithDetail("user id not found").Wrap())
			return
		}
		req.UpdatedBy = userID
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		err = errs.ErrArgs.WrapMsg("invalid json payload " + http.StatusText(http.StatusBadRequest))
		apiresp.GinError(c, err)
		return
	}

	var carousel *entity.DiscoverCarousels
	carousel, err = h.discoverCarouselsUsecase.Edit(ctx, &req)
	if err != nil {
		apiresp.GinError(c, err)
		return
	}
	apiresp.GinSuccess(c, carousel)
}
