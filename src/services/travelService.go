package services

import (
	"api-viajei/src/dto"
	"api-viajei/src/models"
	"api-viajei/src/repository"
	"strconv"
)

func CreateTravel(travelReq dto.TravelDTO) (uint, error) {
	travelId, err := repository.CreateTravel(travelReq)

	if err != nil {
		return 0, err
	}

	return travelId, nil
}

func GetTravels(price string, travelType string) ([]models.Travel, error) {
	priceConverted, err := strconv.ParseFloat(price, 8)

	if err != nil {
		return []models.Travel{}, err
	}

	result, err := repository.GetTravels(priceConverted, travelType)

	if err != nil {
		return []models.Travel{}, err
	}

	return result, nil
}
