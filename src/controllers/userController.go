package controllers

import "github.com/gin-gonic/gin"

func UserController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
