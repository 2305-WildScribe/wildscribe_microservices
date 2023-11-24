package response

import "wildscribe.com/wildscribe/pkg/model"

type AdventureResponse struct {
	Data struct {
		Message    string             `json:"message,omitempty"`
		Error      string             `json:"error,omitempty"`
		Type       string             `json:"type,omitempty"`
		Attributes []*model.Adventure `json:"attributes,omitempty"`
	} `json:"data"`
}
