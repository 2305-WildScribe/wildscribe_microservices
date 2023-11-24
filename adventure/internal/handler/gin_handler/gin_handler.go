package gin_handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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

// GetAnAdventure handles GET /adventure requests.
func (h *Handler) GetAnAdventure() gin.HandlerFunc {
	return func(c *gin.Context) {
		var adventureRequest request.AdventureRequest

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Binds http request to requestBody, if request struct isn't valid, send error
		if err := c.ShouldBindJSON(&adventureRequest); err != nil {
			log.Println(fmt.Errorf("Handler::Show: Failed to Bind JSON to request: %w", err))
			adventureResponse := response.NewErrorResponse("Adventure not found")
			c.JSON(http.StatusBadRequest, adventureResponse)
			return
		}
		// Sends request to controller
		adventure, err := h.ctrl.Show(ctx, adventureRequest)
		if err != nil {
			log.Println(fmt.Errorf("Handler::Show: Error fetching adventure: %w", err))
			adventureResponse := response.NewErrorResponse("Adventure not found")
			c.JSON(http.StatusBadRequest, adventureResponse)
			return
		}

		adventureResponse := response.NewSuccessResponse(adventure)
		c.JSON(http.StatusOK, adventureResponse)
	}
}

// Create Adventure handles Create /adventure requests.
func (h *Handler) CreateAdventure() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var adventureRequest request.AdventureRequest

		// Binds http request to requestBody
		if err := c.ShouldBindJSON(&adventureRequest); err != nil {
			log.Println(fmt.Errorf("Handler::Create: Failed to bind JSON to request: %w", err))
			adventureResponse := response.NewErrorResponse("Adventure Not Created")
			c.JSON(http.StatusBadRequest, adventureResponse)
			return
		}

		adventure, err := h.ctrl.Create(ctx, adventureRequest)
		if err != nil {
			log.Println(fmt.Errorf("Handler::Create: Failed to create adventure: %w", err))
			adventureResponse := response.NewErrorResponse("Adventure Not Created")
			c.JSON(http.StatusUnauthorized, adventureResponse)
			return
		}

		adventureResponse := response.NewSuccessResponse(adventure)
		c.JSON(http.StatusOK, adventureResponse)
	}
}
