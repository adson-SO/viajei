package migrations

import (
	"api-viajei/src/database"
	"api-viajei/src/models"
)

func SyncDatabase() {
	database.DB.AutoMigrate(&models.User{}, &models.Travel{})
}
