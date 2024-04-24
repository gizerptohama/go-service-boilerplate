package database

import (
	"boilerplate/internal/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	entities := []any{
		&entity.User{},
		&entity.ForgotPassword{},
	}

	db.Migrator().DropTable(entities...)

	db.AutoMigrate(entities...)
}
