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
	// User Routes
	router.POST("api/v0/user", handler.LoginUser())
	router.POST("api/v0/user/create", handler.CreateUser())
	router.POST("api/v0/user/update", handler.UpdateUser())
	router.POST("api/v0/user/delete", handler.DeleteUser())

	// Adventure Routes
	router.POST("api/v0/get_adventure", handler.GetAnAdventure())
	router.POST("api/v0/adventures", handler.GetAllAdventures())
	router.POST("api/v0/adventure", handler.CreateAdventure())
	router.PATCH("api/v0/adventure", handler.UpdateAdventure())
	router.DELETE("api/v0/adventure", handler.DeleteAdventure())
}
