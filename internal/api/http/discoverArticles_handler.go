package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/1nterdigital/aka-im-discover/internal/domain"
	"github.com/1nterdigital/aka-im-tools/apiresp"
	"github.com/1nterdigital/aka-im-tools/errs"
	"github.com/1nterdigital/aka-im-tools/log"
	"github.com/1nterdigital/aka-im-tools/mcontext"
	"github.com/1nterdigital/aka-im-tools/tracer"
)

// CreateArticle create an article
//
// @Summary Create a new article
// @Description Creates a new article with the given JSON payload
// @Tags DiscoverArticles
// @Accept json
// @Produce json
// @Param request body domain.DiscoverArticlesAddReq true "Article request"
// @Success 200 {object} domain.DiscoverArticles "Created article"
// @Failure 400 {object} apiresp.ApiResponse "Invalid json payload bad request"
// @Failure 500 {object} apiresp.ApiResponse "Internal server error"
// @Router /bo/discover/articles/add [post]
// @Security ApiKeyAuth
func (h *DiscoverHandler) CreateArticle(c *gin.Context) {
	var (
		req domain.DiscoverArticlesAddReq
		err error
	)

	ctx, span := otel.Tracer(domain.TracerLevelHandler).
		Start(c.Request.Context(), tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.ZError(ctx, "an error occurred while CreateArticle", err)
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

	err = c.ShouldBindJSON(&req)
	if err != nil {
		err = errs.ErrArgs.WrapMsg("invalid json payload " + http.StatusText(http.StatusBadRequest))
		apiresp.GinError(c, err)
		return
	}

	article, err := h.discoverArticlesUsecase.Create(ctx, &req)
	if err != nil {
		apiresp.GinError(c, err)
		return
	}
	apiresp.GinSuccess(c, article)
}

// FindArticles Get paginated list of articles
//
// @Summary Get paginated list of articles
// @Description Retrieves a paginated list of articles
// @Tags DiscoverArticles
// @Accept json
// @Produce json
// @Param id query int false "article id" default("0")
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Page size" default(10)
// @Param title query string false "Title name" default("")
// @Param sortBy query string false "Sort by" default("position")
// @Param order query string false "Order by" default("ASC")
// @Success 200 {array} domain.DiscoverArticles "List of articles"
// @Failure 400 {object} apiresp.ApiResponse "Invalid pagination parameters"
// @Failure 500 {object} apiresp.ApiResponse "Internal server error"
// @Router /bo/discover/articles/find [get]
// @Security ApiKeyAuth
func (h *DiscoverHandler) FindArticles(c *gin.Context) {
	var err error
	ctx, span := otel.Tracer(domain.TracerLevelHandler).
		Start(c.Request.Context(), tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.ZError(ctx, "an error occurred while FindArticles", err)
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

	req := domain.DiscoverArticlesFindReq{
		ID:     id,
		Page:   page,
		Limit:  limit,
		Title:  titleStr,
		SortBy: sortByStr,
		Order:  orderByStr,
	}

	if req.Page <= 0 || req.Limit <= 0 {
		err = errs.ErrArgs.WrapMsg("invalid pagination number: " + http.StatusText(http.StatusBadRequest))
		apiresp.GinError(c, err)
		return
	}

	articles, total, err := h.discoverArticlesUsecase.Find(ctx, &req)
	if err != nil {
		apiresp.GinError(c, err)
		return
	}

	apiresp.GinSuccess(c, gin.H{
		"total": total,
		"data":  articles,
	})
}

// DeleteArticle Delete an article
//
// @Summary Delete an article
// @Description Deletes an article by its ID
// @Tags DiscoverArticles
// @Accept json
// @Produce json
// @Param request body domain.DiscoverArticlesDeleteReq true "Delete request"
// @Success 200 {string} string "deleted"
// @Failure 400 {object} apiresp.ApiResponse "Invalid json payload bad request"
// @Failure 500 {object} apiresp.ApiResponse "Internal server error"
// @Router /discover/articles/delete [delete]
// @Security ApiKeyAuth
func (h *DiscoverHandler) DeleteArticle(c *gin.Context) {
	var (
		req domain.DiscoverArticlesDeleteReq
		err error
	)

	ctx, span := otel.Tracer(domain.TracerLevelHandler).
		Start(c.Request.Context(), tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.ZError(ctx, "an error occurred while DeleteArticle", err)
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

	if err = h.discoverArticlesUsecase.Delete(ctx, req.ID, req.DeletedBy); err != nil {
		apiresp.GinError(c, err)
		return
	}

	apiresp.GinSuccess(c, "deleted")
}

// EditArticle Edit an article
//
// @Summary Edit an article
// @Description Updates an existing article
// @Tags DiscoverArticles
// @Accept json
// @Produce json
// @Param request body domain.DiscoverArticlesEditReq true "Edit request"
// @Success 200 {string} string "updated article"
// @Failure 400 {object} apiresp.ApiResponse "Invalid request payload"
// @Failure 500 {object} apiresp.ApiResponse "Internal server error"
// @Router /discover/articles/edit [post]
// @Security ApiKeyAuth
func (h *DiscoverHandler) EditArticle(c *gin.Context) {
	var (
		req domain.DiscoverArticlesEditReq
		err error
	)

	ctx, span := otel.Tracer(domain.TracerLevelHandler).
		Start(c.Request.Context(), tracer.GetFullFunctionPath())
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			log.ZError(ctx, "an error occurred while EditArticle", err)
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

	err = c.ShouldBindJSON(&req)
	if err != nil {
		err = errs.ErrArgs.WrapMsg("invalid json payload " + http.StatusText(http.StatusBadRequest))
		apiresp.GinError(c, err)
		return
	}

	article, err := h.discoverArticlesUsecase.Edit(ctx, &req)
	if err != nil {
		apiresp.GinError(c, err)
		return
	}
	apiresp.GinSuccess(c, article)
}

func parsePaginationParams(c *gin.Context) (page, limit int32, err error) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	if pageStr == "" {
		pageStr = "10"
	}

	pageInt, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		err = errs.ErrArgs.WrapMsg("invalid page query param")
		return
	}

	if limitStr == "" {
		limitStr = "10"
	}

	limitInt, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		err = errs.ErrArgs.WrapMsg("invalid limit query param")
		return
	}

	return int32(pageInt), int32(limitInt), nil
}

func parseSortOrderParams(c *gin.Context) (sortBy, orderBy string, err error) {
	sortByStr := strings.ToLower(c.DefaultQuery("sortBy", "position"))
	orderByStr := strings.ToUpper(c.DefaultQuery("order", "ASC"))

	if sortByStr == "" {
		sortByStr = "position"
	}

	if sortByStr != "position" && sortByStr != "created_at" {
		err = errs.ErrArgs.WrapMsg("invalid order query param: must be position or created_at")
		return
	}

	if orderByStr == "" {
		orderByStr = "ASC"
	}

	if orderByStr != "ASC" && orderByStr != "DESC" {
		err = errs.ErrArgs.WrapMsg("invalid order query param: must be ASC or DESC")
		return
	}

	return sortByStr, orderByStr, nil
}

func parseIDParam(rawID string) (id int64, err error) {
	if rawID == "" {
		rawID = "0"
	}

	idInt, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		return 0, errs.ErrArgs.WrapMsg("invalid id query param")
	}

	if idInt < 0 {
		return 0, errs.ErrArgs.WrapMsg("invalid id query param")
	}

	return idInt, nil
}

func getOperatedByUser(c *gin.Context, def string) (operatedBy string, err error) {
	if def != "" {
		return def, nil
	}

	userID := mcontext.GetOpUserID(c)
	if userID == "" {
		err = errs.ErrUserIDEmpty.WithDetail("user id not found").Wrap()
		return operatedBy, err
	}

	return userID, nil
}
