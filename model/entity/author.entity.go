package entity

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Country   string         `json:"country"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
