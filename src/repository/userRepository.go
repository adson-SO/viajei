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

func FindUser(email string) (models.User, error) {
	var user models.User
	var err error
	database.DB.First(&user, "email = ?", email)

	if user.ID == 0 {
		return models.User{}, err
	}

	return user, nil
}

func FindUserById(id int64) models.User {
	var user models.User
	database.DB.First(&user, id)

	return user
}
