package request

import (
	"wildscribe.com/user/pkg/model"
)

type UserRequest struct {
	Data struct {
		Type       string `json:"type" binding:"required"`
		Attributes UserAttributes `json:"attributes" binding:"required"`
	} `json:"data" binding:"required"`
}

type UserAttributes struct {
	User_id              string  `json:"user_id,omitempty"`
	Name                 string  `json:"name,omitempty"`
	Email                string  `json:"email,omitempty"`
	Password             string  `json:"password,omitempty"`
}

func (a *UserAttributes) ToUser() *model.User {
	// Convert Attributes to user model
	user := &model.User{
		User_id:              a.User_id,
		Name:                 a.Name,
		Email:                a.Email,
		Password:             a.Password,
	}

	return user
}
