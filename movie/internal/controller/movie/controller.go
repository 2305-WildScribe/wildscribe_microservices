package movie

import (
	"context"
	"errors"

	adventuremodel "movieexample.com/adventure/pkg/model"
	"movieexample.com/movie/internal/gateway"
	"movieexample.com/movie/pkg/model"
	ratingmodel "movieexample.com/rating/pkg/model"
)

var ErrNotFound = errors.New("not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error)
	PutRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType, rating *ratingmodel.Rating) error
}

type adventureGateway interface {
	Get(ctx context.Context, id string) (*adventuremodel.Adventure, error)
}

// Controller defines a movie service controller.
type Controller struct {
	ratingGateway   ratingGateway
	adventureGateway adventureGateway
}

// New creates a new movie service controller.
func New(ratingGateway ratingGateway, adventureGateway adventureGateway) *Controller {
	return &Controller{ratingGateway, adventureGateway}
}

// Get returns movie details including aggregated rating and movie adventure for a given movie ID.
func (c *Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	adventure, err := c.adventureGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	details := &model.MovieDetails{Adventure: *adventure}
	rating, err := c.ratingGateway.GetAggregatedRating(ctx, ratingmodel.RecordID(id), ratingmodel.RecordTypeMovie)
	if err != nil && !errors.Is(err, gateway.ErrNotFound) {
		// Just proceed in this case, it's ok not to have ratings yet.
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}
	return details, nil
}