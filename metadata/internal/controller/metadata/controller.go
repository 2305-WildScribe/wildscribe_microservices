package metadata

import (
	"errors"
	"context"
	"movieexample.com/metadata/pkg/model"
	"movieexample.com/metadata/internal/repository"
	)

	var ErrNotFound = errors.New("not found")

	type metadataRepository interface {
		Get(ctx context.Context, id string) (*model.Metadata, error)
	}

	// Conrtoller defines a metadata service controller.
	type Controller struct {
		repo metadataRepository
	}

	// New creates a new metadata service controller.
	func New(repo metadataRepository) *Controller {
		return &Controller{repo}
	}

	// Get returns movie metadata for a given movie ID.
	func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
		res, err := c.repo.Get(ctx, id)
		if err != nil && errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return res, err
	}