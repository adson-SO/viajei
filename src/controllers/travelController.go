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
			"message": "Error getting data from request body",
		})

		return
	}

	result, err := services.CreateTravel(travel)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating travel in database",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Travel created with id: " + strconv.FormatUint(uint64(result), 10),
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

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"travels": result,
	})
}

func GetTravelsById(c *gin.Context) {
	id := c.Param("id")

	result, err, resultString := services.GetTravelsById(id)

	if err != nil {
		if resultString != "" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Not Found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"travels": result,
	})
}
