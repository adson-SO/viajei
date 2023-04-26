package services

import (
	"api-viajei/src/dto"
	"api-viajei/src/models"
	"api-viajei/src/repository"
	"api-viajei/src/utils"

	"golang.org/x/crypto/bcrypt"
)

func Signup(user dto.UserSignReq) (string, uint, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return "", 0, err
	}

	userID, err := repository.Signup(user.Email, hash)
	if err != nil {
		return "", 0, err
	}

	tokenString, err := utils.GenerateToken(userID)
	if err != nil {
		return "", 0, err
	}

	return tokenString, userID, nil
}

func Signin(userReq dto.UserSignReq) (string, uint, error) {
	email := userReq.Email
	user, err := FindUser(email)
	if err != nil {
		return "", 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
	if err != nil {
		return "", 0, err
	}

	tokenString, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", 0, err
	}

	return tokenString, user.ID, nil
}

func FindUser(email string) (models.User, error) {
	user, err := repository.FindUser(email)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func FindUserById(id float64) (models.User, error) {
	user, err := repository.FindUserById(int64(id))

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
