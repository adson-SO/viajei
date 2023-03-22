package routes

import (
	"api-viajei/src/controllers"
	middleware "api-viajei/src/middlewares"

	"github.com/gin-gonic/gin"
)

func LoadRouter() *gin.Engine {
	router := gin.Default()

	router.Group("api/v1")
	{
		router.POST("/signup", controllers.Signup)
		router.POST("/signin", controllers.Signin)
		router.POST("/travel", middleware.Auth, controllers.CreateTravel)
		router.GET("/travel", middleware.Auth, controllers.GetTravels)
	}

	return router
}
