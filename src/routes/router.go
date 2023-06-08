package routes

import (
	"api-viajei/src/controllers"
	middleware "api-viajei/src/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func LoadRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("api/v1")
	{
		v1.POST("/signup", controllers.Signup)
		v1.POST("/signin", controllers.Signin)
		v1.GET("/user/:id", middleware.Auth, controllers.FindUserById)
		v1.POST("/travel", middleware.Auth, controllers.CreateTravel)
		v1.GET("/travel", middleware.Auth, controllers.GetTravels)
		v1.GET("/travel/:id", middleware.Auth, controllers.GetTravelsById)
		v1.PUT("/user/reset-password", middleware.Auth, controllers.ResetPassword)
	}

	return router
}
