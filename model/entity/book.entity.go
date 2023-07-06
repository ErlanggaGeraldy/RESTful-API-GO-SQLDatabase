package entity

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Published string         `json:"published"`
	Isbn      string         `json:"isbn"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
