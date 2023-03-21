package models

import (
	"time"

	"gorm.io/gorm"
)

type Travel struct {
	gorm.Model
	Destination string    `json:"destination"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Price       float64   `json:"price"`
	Type        string    `json:"type"`
	UserID      uint
}
