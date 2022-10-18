package services

import (
	"api-viajei/src/dto"
	"api-viajei/src/models"
	"api-viajei/src/repository"
	"api-viajei/src/utils"

	"golang.org/x/crypto/bcrypt"
)

func Signup(user dto.UserSignReq) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return "", err
	}

	userID, err := repository.Signup(user.Email, hash)
	if err != nil {
		return "", err
	}

	tokenString, err := utils.GenerateToken(userID)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Signin(userReq dto.UserSignReq) (string, error) {
	email := userReq.Email
	user, err := FindUser(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
	if err != nil {
		return "", err
	}

	tokenString, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func FindUser(email string) (models.User, error) {
	user, err := repository.FindUser(email)

	if err != nil {
		return user, err
	}

	return user, nil
}

func FindUserById(id int64) models.User {
	user := repository.FindUserById(id)

	return user
}
