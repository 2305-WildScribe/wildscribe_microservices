package gin_handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"wildscribe.com/wildscribe/internal/controller/user"
	"wildscribe.com/wildscribe/internal/request"
	"wildscribe.com/wildscribe/internal/response"
)

// Handler defines a movie metadata HTTP handler.
type UserHandler struct {
	ctrl *user.Controller
}

// New creates a new movie metadata HTTP handler.
func NewUserHandler(ctrl *user.Controller) *UserHandler {
	return &UserHandler{ctrl}
}

// GetUser handles GET /user requests.
func (h *UserHandler) GetUser() gin.HandlerFunc {
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
			userResponse.Data.Attributes = map[string]interface{}{"email": userRequest.Data.Attributes.Email, "password": userRequest.Data.Attributes.Password}
			c.JSON(http.StatusUnauthorized, userResponse)
			return
		}

		userResponse.Data.Type = "user"
		userResponse.Data.Attributes = map[string]interface{}{"name": user.Name, "user_id": user.User_id}
		c.JSON(http.StatusOK, userResponse)
	}
}
