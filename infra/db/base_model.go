package db

import (
	"gorm.io/gorm"
	"time"
)

// gorm.Model definition
type Model struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
