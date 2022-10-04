package controllers

import (
	"api-viajei/src/dto"
	"api-viajei/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	user := dto.UserSignupReq{}

	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})

		return
	}

	services.Signup(user)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
