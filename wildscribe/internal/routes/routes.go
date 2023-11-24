package routes

import (
	"github.com/gin-gonic/gin"
	"wildscribe.com/wildscribe/internal/handler/gin_handler"
)

func Routes(router *gin.Engine, handler *gin_handler.GinHandler) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "WILDSCRIBE 2305",
		})
	})
	router.POST("api/v0/user", handler.GetUser())
	router.POST("api/v0/adventure", handler.GetAnAdventure())
}
