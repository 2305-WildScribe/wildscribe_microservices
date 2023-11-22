package adventure

import (
	"context"
	"errors"
	"wildscribe.com/adventure/internal/repository"
	"wildscribe.com/adventure/pkg/model"
)

var ErrNotFound = errors.New("not found")

type adventureRepository interface {
	GetOne(ctx context.Context, id string) (*model.Adventure, error)
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
func (c *Controller) Show(ctx context.Context, id string) (*model.Adventure, error) {
	res, err := c.repo.GetOne(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}
