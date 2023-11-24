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

func NewAdventure(request request.AdventureRequest) *Adventure {
	attributes := request.GetAttributes()

	return &Adventure{
		User_id:              attributes.User_id,
		Adventure_id:         attributes.Adventure_id,
		Activity:             attributes.Activity,
		Date:                 attributes.Date,
		Image_url:            attributes.Image_url,
		Stress_level:         attributes.Stress_level,
		Hours_slept:          attributes.Hours_slept,
		Sleep_stress_notes:   attributes.Sleep_stress_notes,
		Hydration:            attributes.Hydration,
		Diet:                 attributes.Diet,
		Diet_hydration_notes: attributes.Diet_hydration_notes,
		Beta_notes:           attributes.Beta_notes,
	}
}
