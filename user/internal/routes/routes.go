package routes

import (
	"github.com/gin-gonic/gin"
	"wildscribe.com/user/internal/handler/gin_handler"
)

func UserRoute(router *gin.Engine, handler *gin_handler.Handler) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "WILDSCRIBE 2305",
		})
	})
	router.POST("/user", handler.GetUser())
}
