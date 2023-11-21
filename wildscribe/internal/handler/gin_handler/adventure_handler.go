package gin_handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wildscribe.com/wildscribe/internal/controller/adventure"
	"wildscribe.com/wildscribe/internal/request"
	"wildscribe.com/wildscribe/internal/response"
)

// Handler defines a movie metadata HTTP handler.
type AdvHandler struct {
	ctrl *adventure.Controller
}

// New creates a new adventure HTTP handler.
func NewAdvHandler(ctrl *adventure.Controller) *AdvHandler {
	return &AdvHandler{ctrl}
}

// GetAdventure handles POST /adventure requests.
func (h *AdvHandler) GetAnAdventure() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var adventureRequest request.AdventureRequest
		var adventureResponse response.AdventureResponse

		// Binds http request to requestBody
		if err := c.ShouldBindJSON(&adventureRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		adventure, err := h.ctrl.GetAdventure(ctx, adventureRequest)
		if err != nil {
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
