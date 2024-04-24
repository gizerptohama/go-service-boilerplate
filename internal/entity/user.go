package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"column:id;primaryKey;autoIncrement"`
	UserName  string         `gorm:"column:username;type:varchar;not null;check:length(username) > 3;check:length(username) < 10"`
	Email     string         `gorm:"column:email;type:varchar;unique;not null"`
	Password  string         `gorm:"column:password;type:varchar;not null"`
	IsAdmin   bool           `gorm:"column:is_admin;type:bool;not null;default:false"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime:milli;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime:milli;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
