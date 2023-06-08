package routes

import (
	"api-viajei/src/controllers"
	middleware "api-viajei/src/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func LoadRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Authorization", "Content-Type"},
	}))

	router.POST("api/v1/signup", middleware.Cors, controllers.Signup)
	router.POST("api/v1/signin", middleware.Cors, controllers.Signin)
	router.GET("api/v1/user/:id", middleware.Cors, middleware.Auth, controllers.FindUserById)
	router.POST("api/v1/travel", middleware.Cors, middleware.Auth, controllers.CreateTravel)
	router.GET("api/v1/travel", middleware.Cors, middleware.Auth, controllers.GetTravels)
	router.GET("api/v1/travel/:id", middleware.Cors, middleware.Auth, controllers.GetTravelsById)
	router.PUT("api/v1/user/reset-password", middleware.Cors, middleware.Auth, controllers.ResetPassword)

	return router
}
