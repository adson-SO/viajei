package main

import (
	"api-viajei/src/configuration"

	"github.com/gin-gonic/gin"
)

func init() {
	configuration.LoadEnvVariables()
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
