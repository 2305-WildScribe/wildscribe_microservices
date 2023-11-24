package response

import "wildscribe.com/adventure/pkg/model"

type AdventureResponse struct {
	Data struct {
		Message    string             `json:"message,omitempty"`
		Error      string             `json:"error,omitempty"`
		Type       string             `json:"type,omitempty"`
		Attributes []*model.Adventure `json:"attributes,omitempty"`
	} `json:"data"`
}

func NewErrorResponse(err string) *AdventureResponse {
	return &AdventureResponse{
		Data: struct {
			Message    string             `json:"message,omitempty"`
			Error      string             `json:"error,omitempty"`
			Type       string             `json:"type,omitempty"`
			Attributes []*model.Adventure `json:"attributes,omitempty"`
		}{
			Type:  "error",
			Error: err,
		},
	}
}

func NewSuccessResponse(adventures ...*model.Adventure) *AdventureResponse {
	return &AdventureResponse{
		Data: struct {
			Message    string             `json:"message,omitempty"`
			Error      string             `json:"error,omitempty"`
			Type       string             `json:"type,omitempty"`
			Attributes []*model.Adventure `json:"attributes,omitempty"`
		}{
			Type:       "adventure",
			Attributes: adventures,
		},
	}
}
