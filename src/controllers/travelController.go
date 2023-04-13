package controllers

import (
	"api-viajei/src/dto"
	"api-viajei/src/services"
	"net/http"

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

	result, err := services.GetTravels(price, travelType)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"travels": result,
	})
}

func GetTravelsById(c *gin.Context) {
	id := c.Param("id")

	result, err := services.GetTravelsById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"travels": result,
	})
}
