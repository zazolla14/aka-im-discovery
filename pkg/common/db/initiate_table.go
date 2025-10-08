package db

import (
	"fmt"

	"gorm.io/gorm"

	entity "github.com/1nterdigital/aka-im-discover/internal/model"
)

func InitiateTable(gormDB *gorm.DB) error {
	models := []interface{}{
		&entity.DiscoverCarousels{},
		&entity.DiscoverArticles{},
	}
	for _, model := range models {
		if !gormDB.Migrator().HasTable(model) {
			err := gormDB.AutoMigrate(model)
			if err != nil {
				return fmt.Errorf("failed to migrate database: %w", err)
			}
		}
	}

	return nil
}
