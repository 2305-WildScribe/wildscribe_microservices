package controller

import (
	"context"
	"errors"
	"fmt"

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
func (c *Controller) Show(ctx context.Context, adventure_id string) (*model.Adventure, error) {
	adventure, err := c.repo.GetOne(ctx, adventure_id)
	if err != nil {
		new_error := fmt.Errorf("Controller::Show: Failed querying DB: %w", err)
		return nil, new_error
	}
	return adventure, nil
}

func (c *Controller) Index(ctx context.Context, user_id string) ([]*model.Adventure, error) {
	adventures, err := c.repo.GetAll(ctx, user_id)
	if err != nil {
		new_error := fmt.Errorf("Controller::Index: Failed querying DB: %w", err)
		return nil, new_error
	}
	return adventures, nil
}

func (c *Controller) Create(ctx context.Context, adventure *model.Adventure) (*model.Adventure, error) {
	err := c.repo.Create(ctx, adventure)
	if err != nil {
		new_error := fmt.Errorf("Controller::Create: Failed to create adventure: %w", err)
		return nil, new_error
	}
	return adventure, nil
}

func (c *Controller) Update(ctx context.Context, adventure *model.Adventure) (*model.Adventure, error) {
	err := c.repo.Update(ctx, adventure)
	if err != nil {
		new_error := fmt.Errorf("Controller::Create: Failed to update adventure: %w", err)
		return nil, new_error
	}
	return adventure, nil
}

func (c *Controller) Delete(ctx context.Context, adventure_id string) error {
	err := c.repo.Delete(ctx, adventure_id)
	if err != nil {
		new_error := fmt.Errorf("Controller::Create: Failed to update adventure: %w", err)
		return new_error
	}
	return nil
}
