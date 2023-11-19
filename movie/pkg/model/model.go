package model

import "movieexample.com/metadata/pkg/model"

// MovieDetails includes movie metadata and its aggregated ratings.
type MovieDetails struct {
	Rating	 	*float64 				`json:"rating,omitempty"`
	Metadata 	model.Metadata 	`json:"metadata"`
}