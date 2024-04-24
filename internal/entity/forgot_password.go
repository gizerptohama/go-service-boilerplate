package entity

import (
	"time"

	"gorm.io/gorm"
)

type ForgotPassword struct {
	ID        uint           `gorm:"column:id;primaryKey;autoIncrement"`
	UserEmail string         `gorm:"column:user_email;type:varchar;not null"`
	Code      string         `gorm:"column:code;type:varchar;not null"`
	ExpiredAt time.Time      `gorm:"column:expired_at;not null"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime:milli;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime:milli;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      User           `gorm:"foreignKey:user_email;references:email"`
}
