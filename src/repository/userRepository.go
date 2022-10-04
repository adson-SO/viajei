package repository

import (
	"api-viajei/src/database"
	"api-viajei/src/models"
)

func Signup(email string, password []byte) error {
	user := models.User{Email: email, Password: string(password)}
	result := database.DB.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
