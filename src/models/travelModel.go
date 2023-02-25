package models

import (
	"time"

	"gorm.io/gorm"
)

type Travel struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	InitialDate time.Time `json:"initialDate"`
	EndDate     time.Time `json:"endDate"`
	Price       float64   `json:"price"`
	Type        string    `json:"type"`
	Address     string    `json:"address"`
	Tours       []Tour    `json:"tours"`
	UserID      uint
}