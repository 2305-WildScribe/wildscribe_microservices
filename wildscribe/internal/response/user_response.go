package response

import (
	"wildscribe.com/user/pkg/model"
)

type UserResponse struct {
	Data struct {
		Message    string      `json:"message,omitempty"`
		Error      string      `json:"error,omitempty"`
		Type       string      `json:"type,omitempty"`
		Attributes *model.User `json:"attributes,omitempty"`
	} `json:"data"`
}

func NewUserErrorResponse(err string) *UserResponse {
	return &UserResponse{
		Data: struct {
			Message    string      `json:"message,omitempty"`
			Error      string      `json:"error,omitempty"` // Change the type to string
			Type       string      `json:"type,omitempty"`
			Attributes *model.User `json:"attributes,omitempty"`
		}{
			Type:  "error",
			Error: err, // Set the error message here
		},
	}
}

func NewUserSuccessResponse(user *model.User) *UserResponse {
	return &UserResponse{
		Data: struct {
			Message    string      `json:"message,omitempty"`
			Error      string      `json:"error,omitempty"` // Change the type to string
			Type       string      `json:"type,omitempty"`
			Attributes *model.User `json:"attributes,omitempty"`
		}{
			Type:       "user",
			Attributes: user,
		},
	}
}
