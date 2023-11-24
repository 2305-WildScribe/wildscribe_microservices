package adventure

import (
	"context"
	"errors"
	"fmt"
	"wildscribe.com/adventure/internal/request"
	"wildscribe.com/adventure/pkg/model"
)

var ErrNotFound = errors.New("not found")

type adventureRepository interface {
	GetOne(ctx context.Context, id string) (*model.Adventure, error)
	GetAll(ctx context.Context, id string) ([]*model.Adventure, error)
	Create(ctx context.Context, adventure *model.Adventure) error
	Update(ctx context.Context, adventure *model.Adventure) error
	Delete(ctx context.Context, id string) error
}

// Conrtoller defines a adventure service controller.
type Controller struct {
	repo adventureRepository
}

// New creates a new adventure service controller.
func New(repo adventureRepository) *Controller {
	return &Controller{repo}
}

// Get returns an adventure by a given adventure ID.
func (c *Controller) Show(ctx context.Context, request request.AdventureRequest) (*model.Adventure, error) {
	id := request.Data.Attributes.Adventure_id
	adventure, err := c.repo.GetOne(ctx, id)
	if err != nil {
		new_error := fmt.Errorf("Controller::Show: Failed querying DB: %w", err)
		return nil, new_error
	}
	return adventure, nil
}

// Create binds request to a model, Adds it to the Database and returns the model if successful
func (c *Controller) Create(ctx context.Context, request request.AdventureRequest) (*model.Adventure, error) {
	adventure := model.NewAdventure(request)
	err := c.repo.Create(ctx, adventure)
	if err != nil {
		new_error := fmt.Errorf("Controller::Create: Failed to create adventure: %w", err)
		return nil, new_error
	}
	return adventure, nil
}
