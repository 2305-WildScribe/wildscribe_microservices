package request

import "wildscribe.com/adventure/pkg/model"

type AdventureRequest struct {
	Data struct {
		Type       string `json:"type,omitempty" binding:"required"`
		Attributes struct {
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
		} `json:"attributes" binding:"required"`
	} `json:"data" binding:"required"`
}

type GatewayAdventureRequest struct {
	Data struct {
		Type       string             `json:"type,omitempty" binding:"required"`
		Attributes []*model.Adventure `json:"attributes" binding:"required"`
	} `json:"data" binding:"required"`
}
