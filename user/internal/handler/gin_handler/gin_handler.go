package gin_handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wildscribe.com/user/internal/controller/user"
	"wildscribe.com/user/internal/request"
	"wildscribe.com/user/internal/response"
)

// Handler defines a movie metadata HTTP handler.
type Handler struct {
	ctrl *user.Controller
}

// New creates a new movie metadata HTTP handler.
func New(ctrl *user.Controller) *Handler {
	return &Handler{ctrl}
}

// GetUser handles GET /user requests.
func (h *Handler) GetUser() gin.HandlerFunc {
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

		user, err := h.ctrl.Get(ctx, userRequest)
		if err != nil {
			userResponse.Data.Error = "Invalid Email / Password"
			userResponse.Data.Type = "user"
			userResponse.Data.Attributes = user
			c.JSON(http.StatusUnauthorized, userResponse)
			return
		}
		userResponse.Data.Type = "user"
		userResponse.Data.Attributes = user
		c.JSON(http.StatusOK, userResponse)
	}
}
