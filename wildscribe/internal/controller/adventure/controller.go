package adventure

import (
	"context"
	"errors"

	"wildscribe.com/adventure/pkg/model"
	"wildscribe.com/wildscribe/internal/request"
)

var ErrNotFound = errors.New("not found")

type adventureGateway interface {
	GetAdventure(ctx context.Context, request request.AdventureRequest) (*model.Adventure, error)
}

// Controller defines a adventure service controller.
type Controller struct {
	adventureGateway adventureGateway
}

// New creates a new adventure service controller.
func New(gateway adventureGateway) *Controller {
	return &Controller{gateway}
}

// Get returns adventure details for a given adventure ID.
func (c *Controller) GetAdventure(ctx context.Context, request request.AdventureRequest) (*model.Adventure, error) {
	var adventure *model.Adventure
	adventure, err := c.adventureGateway.GetAdventure(ctx, request)
	if err != nil {
		return adventure, err
	}
	return adventure, err
}
