//nolint:dupl // similar to other entity
package domain

import (
	"time"
)

type DiscoverArticles struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	ImageURL  string    `json:"imageUrl"`
	LinkURL   string    `json:"linkUrl"`
	IsActive  bool      `json:"isActive"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy string    `json:"updatedBy"`
	DeletedAt time.Time `json:"deletedAt"`
	DeletedBy string    `json:"deletedBy"`
}

type DiscoverArticlesAddReq struct {
	Title     string `json:"title" binding:"required"`
	ImageURL  string `json:"imageUrl" binding:"required"`
	LinkURL   string `json:"linkUrl" binding:"required"`
	CreatedBy string `json:"createdBy"`
	Position  *int   `json:"position"`
}

type DiscoverArticlesDeleteReq struct {
	ID        int64  `json:"id" binding:"required"`
	DeletedBy string `json:"deletedBy"`
}

type DiscoverArticlesEditReq struct {
	ID        int64  `json:"id" binding:"required"`
	Title     string `json:"title"`
	ImageURL  string `json:"imageUrl"`
	LinkURL   string `json:"linkUrl"`
	UpdatedBy string `json:"updatedBy"`
	Position  *int   `json:"position"`
}

type DiscoverArticlesFindReq struct {
	ID     int64  `json:"id"`
	Page   int32  `validate:"min=1"`
	Limit  int32  `validate:"min=1,max=100"`
	Title  string `json:"title"`
	SortBy string `json:"sortBy"`
	Order  string `json:"order"`
}
