package response

import "wildscribe.com/user/pkg/model"

type UserResponse struct {
	Data struct {
		Message    string      `json:"message,omitempty"`
		Error      string      `json:"error,omitempty"`
		Type       string      `json:"type,omitempty"`
		Attributes *model.User `json:"attributes,omitempty"`
	} `json:"data"`
}
