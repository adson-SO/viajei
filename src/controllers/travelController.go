package controllers

import (
	"api-viajei/src/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTravel(c *gin.Context) {
	travel := dto.TravelCreateReq{}
	err := c.Bind(&travel)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Travel Created",
	})
}
