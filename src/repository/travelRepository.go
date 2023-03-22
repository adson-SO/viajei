package repository

import (
	"api-viajei/src/database"
	"api-viajei/src/dto"
	"api-viajei/src/models"
)

func CreateTravel(travelReq dto.TravelDTO) (uint, error) {
	travel := models.Travel{
		Destination: travelReq.Destination,
		Description: travelReq.Description,
		Date:        travelReq.Date,
		Price:       travelReq.Price,
		Type:        travelReq.Type,
	}

	result := database.DB.Create(&travel)

	if result.Error != nil {
		return 0, result.Error
	}

	return travel.ID, nil
}

func GetTravels(price float64, travelType string) ([]models.Travel, error) {
	var travels []models.Travel
	result := database.DB.Where("type = ?", travelType).Where("price BETWEEN 0 AND ?", price).Find(&travels)

	if result.Error != nil {
		return []models.Travel{}, result.Error
	}

	return travels, nil
}
