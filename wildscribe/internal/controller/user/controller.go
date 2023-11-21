package movie

import (
	"context"
	"errors"

	adventuremodel "wildscribe.com/adventure/pkg/model"
	"wildscribe.com/user/internal/request"
	usermodel "wildscribe.com/user/pkg/model"
	"wildscribe.com/wildscribe/internal/gateway"
	"wildscribe.com/wildscribe/pkg/model"
)

var ErrNotFound = errors.New("not found")

type userGateway interface {
	Get(ctx context.Context, id string) (*adventuremodel.Adventure, error)
}

// Controller defines a movie service controller.
type Controller struct {
	userGateway userGateway
}

// New creates a new movie service controller.
func New(ratingGateway userGateway) *Controller {
	return &Controller{ratingGateway}
}

// Get returns movie details including aggregated rating and movie adventure for a given movie ID.
func (c *Controller) Get(ctx context.Context, request request.UserRequest) (*model.User, error) {
	email, password := request.Email, request.Password
	user, err := c.userGateway.GetUser(email, password)
}
