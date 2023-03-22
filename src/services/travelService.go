package services

import (
	"api-viajei/src/dto"
	"api-viajei/src/models"
	"api-viajei/src/repository"
)

func CreateTravel(travelReq dto.TravelDTO) (uint, error) {
	travelId, err := repository.CreateTravel(travelReq)

	if err != nil {
		return 0, err
	}

	return travelId, nil
}

func GetTravels(price float64, travelType string) ([]models.Travel, error) {
	result, err := repository.GetTravels(price, travelType)

	if err != nil {
		return []models.Travel{}, err
	}

	return result, nil
}
