package services

import (
	"api-viajei/src/dto"
	"api-viajei/src/repository"

	"golang.org/x/crypto/bcrypt"
)

func Signup(user dto.UserSignupReq) error {
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
