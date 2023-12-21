package controller

import (
	"context"
	"fmt"

	"wildscribe.com/adventure/pkg/model"
)

// Interface to allow loose coupling of repos
type adventureRepository interface {
	GetOne(ctx context.Context, id string) (*model.Adventure, error)
	GetAll(ctx context.Context, id string) ([]*model.Adventure, error)
	Create(ctx context.Context, adventure *model.Adventure) error
	Update(ctx context.Context, adventure *model.Adventure) (*model.Adventure, error)
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

// Get returns an adventure by a given adventure ID if success
func (c *Controller) Show(ctx context.Context, adventureId string) (*model.Adventure, error) {
	adventure, err := c.repo.GetOne(ctx, adventureId)
	if err != nil {
		newError := fmt.Errorf("Controller::Show: Failed to get adventure: %w", err)
		return nil, newError
	}
	return adventure, nil
}

// Returns a slice of Adventure Models if success
func (c *Controller) Index(ctx context.Context, userId string) ([]*model.Adventure, error) {
	adventures, err := c.repo.GetAll(ctx, userId)
	if err != nil {
		newError := fmt.Errorf("Controller::Index: Failed to get adventures: %w", err)
		return nil, newError
	}
	return adventures, nil
}

// Adds an adventure, returns adventure if success
func (c *Controller) Create(ctx context.Context, adventure *model.Adventure) (*model.Adventure, error) {
	err := c.repo.Create(ctx, adventure)
	if err != nil {
		newError := fmt.Errorf("Controller::Create: Failed to create adventure: %w", err)
		return nil, newError
	}
	return adventure, nil
}

// Updates an adventure, returns adventure if success
func (c *Controller) Update(ctx context.Context, adventure *model.Adventure) (*model.Adventure, error) {
	updatedAdventure, err := c.repo.Update(ctx, adventure)
	if err != nil {
		newError := fmt.Errorf("Controller::Update: Failed to update adventure: %w", err)
		return nil, newError
	}
	return updatedAdventure, nil
}

// Deletes an adventure, returns nil if success
func (c *Controller) Delete(ctx context.Context, adventureId string) (string, error) {
	err := c.repo.Delete(ctx, adventureId)
	if err != nil {
		newError := fmt.Errorf("Controller::Delete: Failed to Delete adventure: %w", err)
		return "Adventure not found", newError
	}
	return adventureId, nil
}
