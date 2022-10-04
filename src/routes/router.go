package routes

import (
	"api-viajei/src/controllers"

	"github.com/gin-gonic/gin"
)

func LoadRouter() *gin.Engine {
	router := gin.Default()

	router.Group("api/v1")
	{
		router.POST("/signup", controllers.Signup)
	}

	return router
}
