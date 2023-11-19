package request

type UserRequest struct {
	Data struct {
		Type       string `json:"type" binding:"required"`
		Attributes struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		} `json:"attributes" binding:"required"`
	} `json:"data" binding:"required"`
}
