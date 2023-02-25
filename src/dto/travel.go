package dto

import (
	"api-viajei/src/models"
	"time"
)

type TravelCreateReq struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	InitialDate time.Time     `json:"initialDate"`
	EndDate     time.Time     `json:"endDate"`
	Price       float64       `json:"price"`
	Type        string        `json:"type"`
	Address     string        `json:"address"`
	Tours       []models.Tour `json:"tours"`
}
