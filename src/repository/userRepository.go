package repository

import (
	"api-viajei/src/database"
	"api-viajei/src/models"
)

func Signup(name string, email string, password []byte) (uint, error) {
	user := models.User{Name: name, Email: email, Password: string(password)}
	result := database.DB.Create(&user)

	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
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

func FindUserById(id int64) (models.User, error) {
	var user models.User
	var err error
	database.DB.First(&user, id)

	if user.ID == 0 {
		return models.User{}, err
	}

	return user, nil
}

func ResetPassword(email string, password string) error {
	var user models.User
	result := database.DB.Model(&user).Where("email = ?", email).Update("password", password)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
