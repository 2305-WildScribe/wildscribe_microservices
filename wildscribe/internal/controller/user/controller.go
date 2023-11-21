package user

import (
	"context"
	"errors"

	"wildscribe.com/user/pkg/model"
	"wildscribe.com/wildscribe/internal/request"
)

var ErrNotFound = errors.New("not found")

type userGateway interface {
	GetUser(ctx context.Context, request request.UserRequest) (*model.User, error)
}

// Controller defines a user service controller.
type Controller struct {
	userGateway userGateway
}

// New creates a new user service controller.
func New(gateway userGateway) *Controller {
	return &Controller{gateway}
}

// Get returns user details including aggregated rating and user adventure for a given user ID.
func (c *Controller) Get(ctx context.Context, request request.UserRequest) (*model.User, error) {
	var user *model.User
	user, err := c.userGateway.GetUser(ctx, request)
	if err != nil {
		return user, err
	}
	return user, err
}
