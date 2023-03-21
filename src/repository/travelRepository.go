package repository

import (
	"api-viajei/src/database"
	"api-viajei/src/dto"
	"api-viajei/src/models"
)

func CreateTravel(travelReq dto.TravelCreateReq) (uint, error) {
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
