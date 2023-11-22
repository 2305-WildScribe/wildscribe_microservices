package request

type AdventureRequest struct {
	Data struct {
		Type       string `json:"type" binding:"required"`
		Attributes struct {
			Adventure_id string `json:"adventure_id" binding:"required"`
		} `json:"attributes" binding:"required"`
	} `json:"data" binding:"required"`
}
