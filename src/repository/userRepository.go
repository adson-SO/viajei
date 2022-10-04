package repository

import (
	"api-viajei/src/database"
	"api-viajei/src/models"
)

func Signup(email string, password []byte) {
	user := models.User{Email: email, Password: string(password)}
	result := database.DB.Create(&user)

	if result.Error != nil {
		return
	}
}
