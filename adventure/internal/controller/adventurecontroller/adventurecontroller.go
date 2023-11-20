package adventurecontroller

import (
	"context"
	"errors"
	"wildscribe.com/adventure/internal/repository"
	"wildscribe.com/adventure/pkg/model"
)

var ErrNotFound = errors.New("not found")

type adventureRepository interface {
	Get(ctx context.Context, id string) (*model.Adventure, error)
}

// Conrtoller defines a adventure service controller.
type Controller struct {
	repo adventureRepository
}

// New creates a new adventure service controller.
func New(repo adventureRepository) *Controller {
	return &Controller{repo}
}

// Get returns movie adventure for a given movie ID.
func (c *Controller) Get(ctx context.Context, id string) (*model.Adventure, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}
