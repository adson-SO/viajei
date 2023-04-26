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

	tokenString, userId, err := services.Signup(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
		"token":   tokenString,
		"userId":  userId,
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

	tokenString, userId, resultError := services.Signin(user)
	if resultError != nil {
		if tokenString == "Not Found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Not Found",
			})

			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in",
		"token":   tokenString,
		"userId":  userId,
	})
}
