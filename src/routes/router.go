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
		router.POST("/signin", middleware.Auth, controllers.Signin)
	}

	return router
}
