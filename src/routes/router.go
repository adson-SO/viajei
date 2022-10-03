package routes

import (
	"api-viajei/src/controllers"

	"github.com/gin-gonic/gin"
)

func LoadRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/", controllers.UserController)

	return router
}
