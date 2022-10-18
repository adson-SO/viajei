package controllers

import (
	"api-viajei/src/dto"
	"api-viajei/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	user := dto.UserSignReq{}

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})

		return
	}

	resultError := services.Signup(user)
	if resultError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func Signin(c *gin.Context) {
	user := dto.UserSignReq{}

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})

		return
	}

	tokenString, resultError := services.Signin(user)
	if resultError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusNoContent, gin.H{})
}