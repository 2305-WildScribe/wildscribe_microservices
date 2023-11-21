package gin_handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"wildscribe.com/adventure/internal/controller/adventure"
	"wildscribe.com/adventure/internal/request"
	"wildscribe.com/adventure/internal/response"
	// "wildscribe.com/adventure/pkg/model"
)

// Handler defines a movie metadata HTTP handler.
type Handler struct {
	ctrl *adventure.Controller
}

// New creates a new movie metadata HTTP handler.
func New(ctrl *adventure.Controller) *Handler {
	return &Handler{ctrl}
}

// GetAnAdventure handles GET /user requests.
func (h *Handler) GetAnAdventure() gin.HandlerFunc {
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

		adventure, err := h.ctrl.Show(ctx, adventureRequest.Data.Attributes.Adventure_id)
		if err != nil {
			adventureResponse.Data.Error = "Invalid Adventure ID"
			adventureResponse.Data.Type = "adventure"
			c.JSON(http.StatusUnauthorized, adventureResponse)
			return
		}
		adventureResponse.Data.Attributes = append(adventureResponse.Data.Attributes, adventure)

		adventureResponse.Data.Type = "adventure"
		c.JSON(http.StatusOK, adventureResponse)
	}
}
