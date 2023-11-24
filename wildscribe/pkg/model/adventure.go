package model

import "wildscribe.com/wildscribe/internal/request"

type Adventure struct {
	User_id              string `json:"user_id" binding:"required"`
	Adventure_id         string `json:"adventure_id,omitempty" bson:"_id,omitempty"`
	Activity             string `json:"activity" binding:"required"`
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

func NewAdventure(request request.GatewayAdventureRequest) *Adventure {
	return &Adventure{
		User_id:              request.Data.Attributes[0].User_id,
		Adventure_id:         request.Data.Attributes[0].Adventure_id,
		Activity:             request.Data.Attributes[0].Activity,
		Date:                 request.Data.Attributes[0].Date,
		Image_url:            request.Data.Attributes[0].Image_url,
		Stress_level:         request.Data.Attributes[0].Stress_level,
		Hours_slept:          request.Data.Attributes[0].Hours_slept,
		Sleep_stress_notes:   request.Data.Attributes[0].Sleep_stress_notes,
		Hydration:            request.Data.Attributes[0].Hydration,
		Diet:                 request.Data.Attributes[0].Diet,
		Diet_hydration_notes: request.Data.Attributes[0].Diet_hydration_notes,
		Beta_notes:           request.Data.Attributes[0].Beta_notes,
	}
}
