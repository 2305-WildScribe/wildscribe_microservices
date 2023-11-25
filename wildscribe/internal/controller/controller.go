package controller

import (
	"context"
	"errors"
	"fmt"
	"log"

	adventuremodel "wildscribe.com/adventure/pkg/model"
	usermodel "wildscribe.com/user/pkg/model"
	"wildscribe.com/wildscribe/internal/request"
)

var ErrNotFound = errors.New("not found")

type adventureGateway interface {
	GetAdventure(ctx context.Context, adventure_id string) (*adventuremodel.Adventure, error)
	GetAllAdventures(ctx context.Context, user_id string) ([]*adventuremodel.Adventure, error)
}

type userGateway interface {
	GetUser(ctx context.Context, request request.UserRequest) (*usermodel.User, error)
}

// Controller defines a adventure service controller.
type Controller struct {
	adventureGateway adventureGateway
	userGateway      userGateway
}

// New creates a new adventure service controller.
func New(advgateway adventureGateway, usergateway userGateway) *Controller {
	return &Controller{advgateway, usergateway}
}

// Get returns adventure details for a given adventure ID.
func (c *Controller) GetAdventure(ctx context.Context, adventure_id string) (*adventuremodel.Adventure, error) {
	var adventure *adventuremodel.Adventure
	adventure, err := c.adventureGateway.GetAdventure(ctx, adventure_id)
	if err != nil {
		new_error := fmt.Errorf("Controller::GetAdventure: Error fetching Adventure: %w", err)
		log.Println(new_error)
		return adventure, err
	}
	return adventure, err
}

func (c *Controller) GetAllAdventures(ctx context.Context, user_id string) ([]*adventuremodel.Adventure, error) {
	// var adventures []*adventuremodel.Adventure
	adventures, err := c.adventureGateway.GetAllAdventures(ctx, user_id)
	if err != nil {
		new_error := fmt.Errorf("Controller::GetAllAdventures: Error fetching Adventures: %w", err)
		log.Println(new_error)
		return adventures, new_error
	}
	return adventures, err
}

// Get returns user details including aggregated rating and user adventure for a given user ID.
func (c *Controller) GetUser(ctx context.Context, request request.UserRequest) (*usermodel.User, error) {
	var user *usermodel.User
	user, err := c.userGateway.GetUser(ctx, request)
	if err != nil {
		new_error := fmt.Errorf("Controller::GetUser: Error fetching User: %w", err)
		log.Println(new_error)
		return user, err
	}
	return user, err
}
