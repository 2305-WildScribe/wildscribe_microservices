package model

type Adventure struct {
	User_id              string  `json:"user_id" binding:"required"`
	Adventure_id         string  `json:"adventure_id,omitempty" bson:"_id,omitempty"`
	Activity             string  `json:"activity" binding:"required"`
	Date                 string  `json:"date,omitempty"`
	Image_url            string  `json:"image_url,omitempty"`
	Stress_level         string  `json:"stress_level,omitempty"`
	Hours_slept          int32   `json:"hours_slept,omitempty"`
	Sleep_stress_notes   string  `json:"sleep_stress_notes,omitempty"`
	Hydration            string  `json:"hydration,omitempty"`
	Diet                 string  `json:"diet,omitempty"`
	Diet_hydration_notes string  `json:"diet_hydration_notes,omitempty"`
	Beta_notes           string  `json:"beta_notes,omitempty"`
	Lat                  float32 `json:"lat,omitempty"`
	Lon                  float32 `json:"lon,omitempty"`
}
