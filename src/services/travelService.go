package services

import (
	"api-viajei/src/dto"
	"api-viajei/src/models"
	"api-viajei/src/repository"
	"regexp"
	"strconv"
)

func CreateTravel(travelReq dto.TravelDTO) (uint, error, bool) {
	var err error
	isDateStringValid := validateDateString(travelReq.Date)

	if !isDateStringValid {
		return 0, err, false
	}

	travelId, err := repository.CreateTravel(travelReq)

	if err != nil {
		return 0, err, false
	}

	return travelId, nil, true
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

	var query = buildQuery(travelType)

	result, err := repository.GetTravels(query, priceFloat)

	if err != nil {
		return []models.Travel{}, err
	}

	return result, nil
}

func GetTravelsById(id string) ([]models.Travel, error, string) {
	var idUint uint = 0

	if id != "" {
		idConverted, err := strconv.ParseUint(id, 10, 32)

		if err != nil {
			return []models.Travel{}, err, ""
		}

		idUint = uint(idConverted)
	}

	result, err := repository.GetTravelsById(idUint)

	if err != nil {
		return []models.Travel{}, err, "Not Found"
	}

	return result, nil, ""
}

func buildQuery(travelType string) models.Travel {
	var travelTypeQuery string = ""

	if travelType != "" {
		travelTypeQuery = travelType
	}

	return models.Travel{Type: travelTypeQuery}
}

func validateDateString(date string) bool {
	regex := regexp.MustCompile(`^((0)*[1-9]|1[0-9]|2[0-9]|3[01])/((0)*[1-9]|1[012])/(19[0-9][0-9]|20[0-2][0-3])$`)

	return regex.MatchString(date)
}
