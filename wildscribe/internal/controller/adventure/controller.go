package adventure

import (
	"context"
	"errors"
	"log"

	// "wildscribe.com/wildscribe/internal/request"
	"wildscribe.com/adventure/pkg/model"
)

var ErrNotFound = errors.New("not found")

type adventureGateway interface {
	GetAdventure(ctx context.Context, adventure_id string) (*model.Adventure, error)
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
func (c *Controller) GetAdventure(ctx context.Context, adventure_id string) (*model.Adventure, error) {
	var adventure *model.Adventure
	adventure, err := c.adventureGateway.GetAdventure(ctx, adventure_id)
	if err != nil {
		return adventure, err
	}
	log.Println(adventure)
	return adventure, err
}
