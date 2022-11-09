package models

import (
	"time"

	"gorm.io/gorm"
)

type Trip struct {
	gorm.Model
	Title       string    `json: "title"`
	Description string    `json: "description"`
	InitialDate time.Time `json: "initialDate"`
	EndDate     time.Time `json: "endDate"`
	UserID      uint
}
