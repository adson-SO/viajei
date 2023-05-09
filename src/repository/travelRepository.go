package repository

import (
	"api-viajei/src/database"
	"api-viajei/src/dto"
	"api-viajei/src/models"

	"gorm.io/gorm"
)

func CreateTravel(travelReq dto.TravelDTO) (uint, error) {
	travel := models.Travel{
		Destination: travelReq.Destination,
		Description: travelReq.Description,
		Date:        travelReq.Date,
		Price:       travelReq.Price,
		Type:        travelReq.Type,
		Photo:       travelReq.Photo,
		UserID:      travelReq.UserID,
	}

	result := database.DB.Create(&travel)

	if result.Error != nil {
		return 0, result.Error
	}

	return travel.ID, nil
}

func GetTravels(query models.Travel, price float64) ([]models.Travel, error) {
	var travels []models.Travel
	var result *gorm.DB

	if price != 0 {
		result = database.DB.Where(&query).Where("price >= 0 AND price <= ?", price).Find(&travels)
	} else {
		result = database.DB.Where(&query).Find(&travels)
	}

	if result.Error != nil {
		return []models.Travel{}, result.Error
	}

	return travels, nil
}

func GetTravelsById(id uint) ([]models.Travel, error) {
	var travels []models.Travel
	result := database.DB.Where(&models.Travel{UserID: id}).Find(&travels)

	if result.Error != nil {
		return []models.Travel{}, result.Error
	}

	return travels, nil
}
