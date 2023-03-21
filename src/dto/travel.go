package dto

import (
	"time"
)

type TravelCreateReq struct {
	Destination string    `json:"destination"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Price       float64   `json:"price"`
	Type        string    `json:"type"`
}
