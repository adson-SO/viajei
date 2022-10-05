package services

import (
	"api-viajei/src/dto"
	"api-viajei/src/models"
	"api-viajei/src/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(user dto.UserSignReq) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return err
	}

	resultError := repository.Signup(user.Email, hash)

	if resultError != nil {
		return resultError
	}

	return nil
}

func Signin(userReq dto.UserSignReq) (string, error) {
	email := userReq.Email
	user, _ := FindUser(email)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
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
