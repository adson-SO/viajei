package routes

import (
	"api-viajei/src/controllers"
	middleware "api-viajei/src/middlewares"

	"github.com/gin-gonic/gin"
)

func LoadRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/signup", middleware.Cors, controllers.Signup)
		v1.POST("/signin", middleware.Cors, controllers.Signin)
		v1.GET("/user/:id", middleware.Cors, middleware.Auth, controllers.FindUserById)
		v1.POST("/travel", middleware.Cors, middleware.Auth, controllers.CreateTravel)
		v1.GET("/travel", middleware.Cors, middleware.Auth, controllers.GetTravels)
		v1.GET("/travel/:id", middleware.Cors, middleware.Auth, controllers.GetTravelsById)
		v1.PUT("/user/reset-password", middleware.Cors, middleware.Auth, controllers.ResetPassword)
	}

	return router
}
