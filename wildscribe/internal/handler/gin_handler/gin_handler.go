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

// LoginUser handles user login requests.
func (h *GinHandler) LoginUser() gin.HandlerFunc {
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
		user := userRequest.Data.Attributes.ToUser()
		resp_user, err := h.ctrl.LoginUser(ctx, user)
		if err != nil {
			new_error := fmt.Errorf("GinHandler::LoginUser: Error fetching User: %w", err)
			log.Println(new_error)
			response := response.NewErrorResponse("Invalid Email/Password")
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		log.Println(resp_user)
		userResponse.Data.Type = "user"
		userResponse.Data.Attributes.User_id = resp_user.User_id
		c.JSON(http.StatusOK, userResponse)
	}
}

func (h *GinHandler) CreateUser() gin.HandlerFunc {
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
		user := userRequest.Data.Attributes.ToUser()
		resp_user, err := h.ctrl.CreateUser(ctx, user)
		if err != nil {
			new_error := fmt.Errorf("GinHandler::CreateUser: Error fetching User: %w", err)
			log.Println(new_error)
			response := response.NewUserErrorResponse("Invalid Email/Password")
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		log.Println(user)
		userResponse.Data.Type = "user"
		userResponse.Data.Attributes.User_id = resp_user.User_id
		c.JSON(http.StatusOK, userResponse)
	}
}

func (h *GinHandler) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var userRequest request.UserRequest
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			new_error := fmt.Errorf("GinHandler::UpdateUser: Error binding request body to JSON: %w", err)
			log.Println(new_error)
			response := response.NewUserErrorResponse("Invalid Request Format")
			c.JSON(http.StatusBadRequest, response)
			return
		}
		user, err := h.ctrl.UpdateUser(ctx, userRequest.Data.Attributes.ToUser())
		if err != nil {
			new_error := fmt.Errorf("GinHandler::UpdateUser: Error fetching User: %w", err)
			log.Println(new_error)
			response := response.NewUserErrorResponse("Invalid Email/Password")
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		response := response.NewUserSuccessResponse(user)
		c.JSON(http.StatusOK, response)
	}
}

func (h *GinHandler) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var userRequest request.UserRequest
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			new_error := fmt.Errorf("GinHandler::DeleteUser: Error binding request body to JSON: %w", err)
			log.Println(new_error)
			response := response.NewUserErrorResponse("Invalid Request Format")
			c.JSON(http.StatusBadRequest, response)
			return
		}
		user, err := h.ctrl.DeleteUser(ctx, userRequest.Data.Attributes.User_id)
		if err != nil {
			new_error := fmt.Errorf("GinHandler::DeleteUser: Error fetching User: %w", err)
			log.Println(new_error)
			response := response.NewUserErrorResponse("Invalid Email/Password")
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		var userResponse response.UserResponse
		userResponse.Data.Type = "user"
		userResponse.Data.Attributes.User_id = user
		c.JSON(http.StatusOK, user)
	}
}
		

// GetAdventure handles POST /adventure requests.
func (h *GinHandler) GetAnAdventure() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var adventureRequest request.AdventureRequest
		// Binds http request to requestBody
		if err := c.ShouldBindJSON(&adventureRequest); err != nil {
			new_error := fmt.Errorf("GinHandler::GetAnAdventure: Error binding request body to JSON: %w", err)
			log.Println(new_error)
			response := response.NewErrorResponse("Invalid Request Format")
			c.JSON(http.StatusBadRequest, response)
			return
		}
		adventure, err := h.ctrl.GetAdventure(ctx, adventureRequest.Data.Attributes.Adventure_id)
		if err != nil {
			new_error := fmt.Errorf("GinHandler::GetAnAdventure: Error fetching Adventure: %w", err)
			log.Println(new_error)
			response := response.NewErrorResponse("Invalid Adventure ID")
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		response := response.NewSuccessResponse(adventure)
		c.JSON(http.StatusOK, response)
	}
}

func (h *GinHandler) GetAllAdventures() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var adventureRequest request.AdventureRequest
		// Binds http request to requestBody
		if err := c.ShouldBindJSON(&adventureRequest); err != nil {
			new_error := fmt.Errorf("GinHandler::GetAllAdventures: Error binding request body to JSON: %w", err)
			log.Println(new_error)
			response := response.NewErrorResponse("Invalid Request Format")
			c.JSON(http.StatusBadRequest, response)
			return
		}
		adventures, err := h.ctrl.GetAllAdventures(ctx, adventureRequest.Data.Attributes.User_id)
		if err != nil {
			new_error := fmt.Errorf("GinHandler::GetAllAdventures: Error fetching Adventure: %w", err)
			log.Println(new_error)
			response := response.NewErrorResponse("Invalid Adventure ID")
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		log.Println(adventures)
		response := response.NewSuccessResponse(adventures...)
		c.JSON(http.StatusOK, response)
	}
}

func (h *GinHandler) CreateAdventure() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var adventureRequest request.AdventureRequest
		// Binds http request to requestBody
		if err := c.ShouldBindJSON(&adventureRequest); err != nil {
			new_error := fmt.Errorf("GinHandler::CreateAdventure: Error binding request body to JSON: %w", err)
			log.Println(new_error)
			response := response.NewErrorResponse("Invalid Request Format")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		adventure := adventureRequest.Data.Attributes.ToAdventure()
		createdAdventure, err := h.ctrl.CreateAdventure(ctx, adventure)
		if err != nil {
			new_error := fmt.Errorf("GinHandler::CreateAdventure: Error fetching Adventure: %w", err)
			log.Println(new_error)
			response := response.NewErrorResponse("Invalid Adventure ID")
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		response := response.NewSuccessResponse(createdAdventure)
		c.JSON(http.StatusOK, response)
	}
}
func (h *GinHandler) UpdateAdventure() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var adventureRequest request.AdventureRequest
		// Binds http request to requestBody
		if err := c.ShouldBindJSON(&adventureRequest); err != nil {
			new_error := fmt.Errorf("GinHandler::GetAnAdventure: Error binding request body to JSON: %w", err)
			log.Println(new_error)
			response := response.NewErrorResponse("Invalid Request Format")
			c.JSON(http.StatusBadRequest, response)
			return
		}
		adventure := adventureRequest.Data.Attributes.ToAdventure()
		createdAdventure, err := h.ctrl.UpdateAdventure(ctx, adventure)
		if err != nil {
			new_error := fmt.Errorf("GinHandler::GetAnAdventure: Error fetching Adventure: %w", err)
			log.Println(new_error)
			response := response.NewErrorResponse("Invalid Adventure ID")
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		response := response.NewSuccessResponse(createdAdventure)
		c.JSON(http.StatusOK, response)
	}
}

func (h *GinHandler) DeleteAdventure() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var adventureRequest request.AdventureRequest
		// Binds http request to requestBody
		if err := c.ShouldBindJSON(&adventureRequest); err != nil {
			new_error := fmt.Errorf("GinHandler::GetAnAdventure: Error binding request body to JSON: %w", err)
			log.Println(new_error)
			response := response.NewErrorResponse("Invalid Request Format")
			c.JSON(http.StatusBadRequest, response)
			return
		}
		_, err := h.ctrl.DeleteAdventure(ctx, adventureRequest.Data.Attributes.Adventure_id)
		if err != nil {
			new_error := fmt.Errorf("GinHandler::GetAnAdventure: Error fetching Adventure: %w", err)
			log.Println(new_error)
			response := response.NewErrorResponse("Invalid Adventure ID")
			c.JSON(http.StatusUnauthorized, response)
			return
		}
		response := response.NewSuccessResponse(nil)
		c.JSON(http.StatusOK, response)
	}
}
