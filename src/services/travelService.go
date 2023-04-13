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
	var priceFloat float64 = 0

	if price != "" {
		priceConverted, err := strconv.ParseFloat(price, 8)

		if err != nil {
			return []models.Travel{}, err
		}

		priceFloat = priceConverted
	}

	var query = buildQuery(priceFloat, travelType)

	result, err := repository.GetTravels(query)

	if err != nil {
		return []models.Travel{}, err
	}

	return result, nil
}

func GetTravelsById(id string) ([]models.Travel, error) {
	var idUint uint = 0

	if id != "" {
		idConverted, err := strconv.ParseUint(id, 10, 32)

		if err != nil {
			return []models.Travel{}, err
		}

		idUint = uint(idConverted)
	}

	result, err := repository.GetTravelsById(idUint)

	if err != nil {
		return []models.Travel{}, err
	}

	return result, nil
}

func buildQuery(price float64, travelType string) models.Travel {
	var priceQuery float64 = 0
	var travelTypeQuery string = ""

	if price != 0 {
		priceQuery = price
	}

	if travelType != "" {
		travelTypeQuery = travelType
	}

	return models.Travel{Price: priceQuery, Type: travelTypeQuery}
}
