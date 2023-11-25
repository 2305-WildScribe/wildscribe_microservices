package request

import (
	"wildscribe.com/adventure/pkg/model"
)

type AdventureRequest struct {
	Data struct {
		Type       string     `json:"type,omitempty" binding:"required"`
		Attributes Attributes `json:"attributes" binding:"required"`
	} `json:"data" binding:"required"`
}

type Attributes struct {
	User_id              string  `json:"user_id,omitempty"`
	Adventure_id         string  `json:"adventure_id,omitempty"`
	Activity             string  `json:"activity,omitempty"`
	Date                 string  `json:"date,omitempty"`
	Image_url            string  `json:"image_url,omitempty"`
	Stress_level         string  `json:"stress_level,omitempty"`
	Hours_slept          int     `json:"hours_slept,omitempty"`
	Sleep_stress_notes   string  `json:"sleep_stress_notes,omitempty"`
	Hydration            string  `json:"hydration,omitempty"`
	Diet                 string  `json:"diet,omitempty"`
	Diet_hydration_notes string  `json:"diet_hydration_notes,omitempty"`
	Beta_notes           string  `json:"beta_notes,omitempty"`
	Lat                  float32 `json:"lat,omitempty"`
	Lon                  float32 `json:"lon,omitempty"`
}

func (a *Attributes) ToAdventure() *model.Adventure {
	// Convert Attributes to Adventure model
	adventure := &model.Adventure{
		User_id:              a.User_id,
		Adventure_id:         a.Adventure_id,
		Activity:             a.Activity,
		Date:                 a.Date,
		Image_url:            a.Image_url,
		Stress_level:         a.Stress_level,
		Hours_slept:          int32(a.Hours_slept),
		Sleep_stress_notes:   a.Sleep_stress_notes,
		Hydration:            a.Hydration,
		Diet:                 a.Diet,
		Diet_hydration_notes: a.Diet_hydration_notes,
		Beta_notes:           a.Beta_notes,
		Lat:                  a.Lat,
		Lon:                  a.Lon,
	}

	return adventure
}
