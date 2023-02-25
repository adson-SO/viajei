package repository

import (
	"api-viajei/src/database"
	"api-viajei/src/dto"
	"api-viajei/src/models"
)

func CreateTravel(travelReq dto.TravelCreateReq) (uint, error) {
	travel := models.Travel{
		Title:       travelReq.Title,
		Description: travelReq.Description,
		InitialDate: travelReq.InitialDate,
		EndDate:     travelReq.EndDate,
		Price:       travelReq.Price,
		Type:        travelReq.Type,
		Address:     travelReq.Address,
		Tours:       travelReq.Tours,
	}

	result := database.DB.Create(&travel)

	if result.Error != nil {
		return 0, result.Error
	}

	return travel.ID, nil
}
