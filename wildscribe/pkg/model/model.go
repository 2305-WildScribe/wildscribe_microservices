package model

import "wildscribe.com/adventure/pkg/model"

// MovieDetails includes movie metadata and its aggregated ratings.
type MovieDetails struct {
	Rating    *float64        `json:"rating,omitempty"`
	Adventure model.Adventure `json:"adventure"`
}
