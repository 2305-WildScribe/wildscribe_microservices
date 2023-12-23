package controller

import (
	"context"
	"errors"
	"fmt"

	"wildscribe.com/user/internal/middleware/bcrypt_middleware"
	"wildscribe.com/user/pkg/model"
)

// ErrNotFound is returned when a record is not found.
var ErrNotFound = errors.New("users not found for a record")

type userRepository interface {
	Get(ctx context.Context, email string) (*model.User, error)
	Validate(ctx context.Context, user_id string) (bool, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user_id string) error
}

// Controller defines a user service controller.
type Controller struct {
	repo userRepository
}

// New creates a new user service controller.
func New(repo userRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) Login(ctx context.Context, email string, password string) (*model.User, error) {
	// Sets email and password variables, password must be []byte for bcrypt use
	byte_password := []byte(password)

	user, err := c.repo.Get(ctx, email)
	if err != nil {
		new_error := fmt.Errorf("Controller::Login: failed to fetch user %w", err)
		return nil, new_error
	}

	if !bcrypt_middleware.ComparePasswords(user.Password, byte_password) {
		new_error := fmt.Errorf("Controller::Login: failed to compare passwords")
		return nil, new_error
	}

	return user, err
}

func (c *Controller) Validate(ctx context.Context, user_id string) (bool, error) {
	user, err := c.repo.Validate(ctx, user_id)
	if err != nil {
		new_error := fmt.Errorf("Controller::Validate: failed to fetch user %w", err)
		return false, new_error
	}
	return user, err
}

func (c *Controller) Create(ctx context.Context, user *model.User) (*model.User, error) {
	// Check if email exist before creating user
	existingUser, check_err := c.repo.Get(ctx, user.Email)
	if check_err == nil && existingUser != nil {
		// User with the same email already exists
		return nil, fmt.Errorf("Controller::Create: email '%s' is already taken", user.Email)
	}
	// Convert password to []byte for bcrypt
	byte_password := []byte(user.Password)
	// Save new HashandSalted password
	user.Password = bcrypt_middleware.HashAndSalt(byte_password)
	err := c.repo.Create(ctx, user)
	if err != nil {
		new_error := fmt.Errorf("Controller::Create: failed to create user %w", err)
		return nil, new_error
	}
	return user, err
}

func (c *Controller) Update(ctx context.Context, user *model.User) (*model.User, error) {
	byte_password := []byte(user.Password)
	user.Password = bcrypt_middleware.HashAndSalt(byte_password)
	err := c.repo.Update(ctx, user)
	if err != nil {
		new_error := fmt.Errorf("Controller::Update: failed to update user %w", err)
		return nil, new_error
	}
	user.Password = ""
	return user, err
}

func (c *Controller) Delete(ctx context.Context, user_id string) error {
	err := c.repo.Delete(ctx, user_id)
	if err != nil {
		new_error := fmt.Errorf("Controller::Delete: failed to delete user %w", err)
		return new_error
	}
	return err
}
