package user

import (
	"context"
	"errors"
	"wildscribe.com/user/internal/middleware/bcrypt_middleware"
	"wildscribe.com/user/internal/request"
	"wildscribe.com/user/pkg/model"
)

// ErrNotFound is returned when a record is not found.
var ErrNotFound = errors.New("users not found for a record")

type userRepository interface {
	Get(ctx context.Context, email string) (*model.User, error)
}

// Controller defines a user service controller.
type Controller struct {
	repo userRepository
}

// New creates a new user service controller.
func New(repo userRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) Get(ctx context.Context, request request.UserRequest) (*model.User, error) {

	// Sets email and password variables, password must be []byte for bcrypt use
	email := request.Data.Attributes.Email
	password := []byte(request.Data.Attributes.Password)

	user, err := c.repo.Get(ctx, email)
	if err != nil {
		return nil, err
	}

	if bcrypt_middleware.ComparePasswords(user.Password, password) == false {
		return nil, err
	}

	return user, err
}
