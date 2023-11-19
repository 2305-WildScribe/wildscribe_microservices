package response

type UserResponse struct {
	Data struct {
		Message    string                 `json:"message,omitempty"`
		Error      string                 `json:"error,omitempty"`
		Type       string                 `json:"type,omitempty"`
		Attributes map[string]interface{} `json:"attributes,omitempty"`
	} `json:"data"`
}
