package memory

import ( 
				"context"
				 "sync"

				 "movieexample.com/metadata/pkg/model"
				 "movieexample.com/metadata/internal/repository"
)
// Defines a memory movie metadata repository.
type Repository struct { 
	sync.RWMutex
	data map[string] *model.Metadata
}

// New creates a new memory repository.
func New() *Repository {
	return &Repository{
		data: map[string] *model.Metadata{},
	}
}

// Get retrieves a movie metadate by movie ID.
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()
	metadata, ok := r.data[id]

	if !ok {
		return nil, repository.ErrNotFound
	}

	return metadata, nil
}

// Puts adds movie metadata for a given movie ID.

func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}