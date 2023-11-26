package controller

import (
	"context"
	"errors"
	"fmt"
	"log"

	adventuremodel "wildscribe.com/adventure/pkg/model"
	usermodel "wildscribe.com/user/pkg/model"
)

var ErrNotFound = errors.New("not found")

type adventureGateway interface {
	GetAdventure(ctx context.Context, adventure_id string) (*adventuremodel.Adventure, error)
	GetAllAdventures(ctx context.Context, user_id string) ([]*adventuremodel.Adventure, error)
	CreateAdventure(ctx context.Context, adventure *adventuremodel.Adventure) (*adventuremodel.Adventure, error)
	UpdateAdventure(ctx context.Context, adventure *adventuremodel.Adventure) (*adventuremodel.Adventure, error)
	DeleteAdventure(ctx context.Context, adventure_id string) (string, error)
}

type userGateway interface {
	LoginUser(ctx context.Context, user *usermodel.User) (*usermodel.User, error)
	ValidateUser(ctx context.Context, user_id string) (string, error)
	CreateUser(ctx context.Context, user *usermodel.User) (*usermodel.User, error)
	UpdateUser(ctx context.Context, user *usermodel.User) (*usermodel.User, error)
	DeleteUser(ctx context.Context, user_id string) (string, error)
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

func (c *Controller) CreateAdventure(ctx context.Context, adventure *adventuremodel.Adventure) (*adventuremodel.Adventure, error) {
	// var adventures []*adventuremodel.Adventure
	adventure, err := c.adventureGateway.CreateAdventure(ctx, adventure)
	if err != nil {
		new_error := fmt.Errorf("Controller::GetAllAdventures: Error fetching Adventures: %w", err)
		log.Println(new_error)
		return adventure, new_error
	}
	return adventure, err
}

func (c *Controller) UpdateAdventure(ctx context.Context, adventure *adventuremodel.Adventure) (*adventuremodel.Adventure, error) {
	// var adventures []*adventuremodel.Adventure
	adventure, err := c.adventureGateway.UpdateAdventure(ctx, adventure)
	if err != nil {
		new_error := fmt.Errorf("Controller::GetAllAdventures: Error fetching Adventures: %w", err)
		log.Println(new_error)
		return adventure, new_error
	}
	return adventure, err
}

func (c *Controller) DeleteAdventure(ctx context.Context, adventure_id string) (string, error) {
	// var adventures []*adventuremodel.Adventure
	adventure, err := c.adventureGateway.DeleteAdventure(ctx, adventure_id)
	if err != nil {
		new_error := fmt.Errorf("Controller::GetAllAdventures: Error fetching Adventures: %w", err)
		log.Println(new_error)
		return adventure, new_error
	}
	return adventure, err
}

// Get returns user details including aggregated rating and user adventure for a given user ID.
func (c *Controller) LoginUser(ctx context.Context, user *usermodel.User) (*usermodel.User, error) {
	user, err := c.userGateway.LoginUser(ctx, user)
	if err != nil {
		new_error := fmt.Errorf("Controller::GetUser: Error fetching User: %w", err)
		log.Println(new_error)
		return user, err
	}
	return user, err
}

func (c *Controller) ValidateUser(ctx context.Context, user_id string) (string, error) {
	var user string
	user, err := c.userGateway.ValidateUser(ctx, user_id)
	if err != nil {
		new_error := fmt.Errorf("Controller::ValidateUser: Error fetching User: %w", err)
		log.Println(new_error)
		return user, err
	}
	return user, err
}

func (c *Controller) CreateUser(ctx context.Context, user *usermodel.User) (*usermodel.User, error) {
	resp_user, err := c.userGateway.CreateUser(ctx, user)
	if err != nil {
		new_error := fmt.Errorf("Controller::CreateUser: Error fetching User: %w", err)
		log.Println(new_error)
		return user, err
	}
	resp_user.Password = ""
	return resp_user, err
}

func (c *Controller) UpdateUser(ctx context.Context, user *usermodel.User) (*usermodel.User, error) {
	user, err := c.userGateway.UpdateUser(ctx, user)
	if err != nil {
		new_error := fmt.Errorf("Controller::UpdateUser: Error fetching User: %w", err)
		log.Println(new_error)
		return user, err
	}
	return user, err
}

func (c *Controller) DeleteUser(ctx context.Context, user_id string) (string, error) {
	user, err := c.userGateway.DeleteUser(ctx, user_id)
	if err != nil {
		new_error := fmt.Errorf("Controller::DeleteUser: Error fetching User: %w", err)
		log.Println(new_error)
		return user, err
	}
	return user, err
}
