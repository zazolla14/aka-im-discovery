package entity

import (
	"time"
)

type DiscoverArticles struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title     string     `gorm:"column:title" json:"title"`
	ImageURL  string     `gorm:"column:image_url" json:"imageUrl"`
	LinkURL   string     `gorm:"column:link_url" json:"linkUrl"`
	IsActive  bool       `gorm:"column:is_active;not null; default:true" json:"isActive"`
	Position  *int       `gorm:"column:position" json:"position"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	CreatedBy string     `gorm:"column:created_by" json:"createdBy"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	UpdatedBy string     `gorm:"column:updated_by" json:"updatedBy"`
	DeletedAt *time.Time `gorm:"column:deleted_at; default:null" json:"deletedAt"`
	DeletedBy string     `gorm:"column:deleted_by" json:"deletedBy"`
}

func (DiscoverArticles) TableName() string {
	return "articles"
}
