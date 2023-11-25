package gin_handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"wildscribe.com/wildscribe/internal/controller"
	"wildscribe.com/wildscribe/internal/request"
	"wildscribe.com/wildscribe/internal/response"
)

// Handler defines a movie metadata HTTP handler.
type GinHandler struct {
	ctrl *controller.Controller
}

// New creates a new movie metadata HTTP handler.
func NewGinHandler(ctrl *controller.Controller) *GinHandler {
	return &GinHandler{ctrl}
}

// GetUser handles GET /user requests.
func (h *GinHandler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var userRequest request.UserRequest
		var userResponse response.UserResponse

		// Binds http request to requestBody
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := h.ctrl.GetUser(ctx, userRequest)
		if err != nil {
			userResponse.Data.Error = "Invalid Email / Password"
			userResponse.Data.Type = "user"
			userResponse.Data.Attributes = nil
			c.JSON(http.StatusUnauthorized, userResponse)
			return
		}
		log.Println(user)
		userResponse.Data.Type = "user"
		userResponse.Data.Attributes = user
		c.JSON(http.StatusOK, userResponse)
	}
}

// GetAdventure handles POST /adventure requests.
func (h *GinHandler) GetAnAdventure() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var adventureRequest request.AdventureRequest
		var adventureResponse response.AdventureResponse

		// Binds http request to requestBody
		if err := c.ShouldBindJSON(&adventureRequest); err != nil {
			log.Println(fmt.Errorf("Gin_Handler::GetAnAdventure: Error Binding Request JSON: %w.", err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		adventure, err := h.ctrl.GetAdventure(ctx, adventureRequest.Data.Attributes.Adventure_id)
		if err != nil {
			log.Println(fmt.Errorf("Gin_Handler::GetAnAdventure: Couldn't Fetch Adventure: %w.", err))
			adventureResponse.Data.Error = "Invalid Email / Password"
			adventureResponse.Data.Type = "adventure"
			adventureResponse.Data.Attributes = nil
			c.JSON(http.StatusUnauthorized, adventureResponse)
			return
		}

		adventureResponse.Data.Type = "adventure"
		adventureResponse.Data.Attributes = append(adventureResponse.Data.Attributes, adventure)
		c.JSON(http.StatusOK, adventureResponse)
	}
}
