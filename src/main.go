package main

import (
	"api-viajei/src/configuration"
	"api-viajei/src/database"

	"github.com/gin-gonic/gin"
)

func init() {
	configuration.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.Run()
}
