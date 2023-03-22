package controllers

import (
	"api-viajei/src/dto"
	"api-viajei/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTravel(c *gin.Context) {
	travel := dto.TravelDTO{}
	err := c.Bind(&travel)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	result, err := services.CreateTravel(travel)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Travel created with id: " + string(result),
	})
}

func GetTravels(c *gin.Context) {
	price := c.Query("price")
	travelType := c.Query("type")

	priceConverted, err := strconv.ParseFloat(price, 8)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	result, err := services.GetTravels(float64(priceConverted), travelType)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"travels": result,
	})
}
