package request

import (
// "wildscribe.com/adventure/internal/request"
// "wildscribe.com/adventure/pkg/model"
)

type AdventureRequest struct {
	Data struct {
		Type       string     `json:"type,omitempty" binding:"required"`
		Attributes Attributes `json:"attributes" binding:"required"`
	} `json:"data" binding:"required"`
}

type Attributes struct {
	User_id              string `json:"user_id,omitempty"`
	Adventure_id         string `json:"adventure_id,omitempty"`
	Activity             string `json:"activity,omitempty"`
	Date                 string `json:"date,omitempty"`
	Image_url            string `json:"image_url,omitempty"`
	Stress_level         string `json:"stress_level,omitempty"`
	Hours_slept          int    `json:"hours_slept,omitempty"`
	Sleep_stress_notes   string `json:"sleep_stress_notes,omitempty"`
	Hydration            string `json:"hydration,omitempty"`
	Diet                 string `json:"diet,omitempty"`
	Diet_hydration_notes string `json:"diet_hydration_notes,omitempty"`
	Beta_notes           string `json:"beta_notes,omitempty"`
}
type AdventureRequestInterface interface {
	GetAttributes() Attributes
}

type GatewayAdventureRequest struct {
	Data struct {
		Type       string     `json:"type,omitempty" binding:"required"`
		Attributes Attributes `json:"attributes" binding:"required"`
	} `json:"data" binding:"required"`
}

type OtherAdventureRequest struct {
	// Define the structure for the other type of request
}

func (r *AdventureRequest) GetAttributes() *Attributes {
	return &r.Data.Attributes
}
