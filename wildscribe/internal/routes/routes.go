package routes

import (
	"github.com/gin-gonic/gin"
	"wildscribe.com/wildscribe/internal/handler/gin_handler"
)

func UserRoutes(router *gin.Engine, handler *gin_handler.UserHandler) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "WILDSCRIBE 2305",
		})
	})
	router.POST("api/v0/user", handler.GetUser())
}

func AdventureRoutes(router *gin.Engine, handler *gin_handler.AdvHandler) {
	router.POST("api/v0/adventure", handler.GetAnAdventure())
}
