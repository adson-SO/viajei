package services

import (
	"api-viajei/src/dto"
	"api-viajei/src/repository"
)

func CreateTravel(travelReq dto.TravelCreateReq) (uint, error) {
	travelId, err := repository.CreateTravel(travelReq)

	if err != nil {
		return 0, err
	}

	return travelId, nil
}
