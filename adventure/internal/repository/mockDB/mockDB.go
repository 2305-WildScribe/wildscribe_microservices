package mockDB

import (
	"context"
	"wildscribe.com/adventure/pkg/model"
)

// Set Collection struct
type Collection struct {
	collection string
}

// Sets the collection to "Test"
func MockCollection() *Collection {
	return &Collection{
		collection: "Test",
	}
}

// Get a single collection from the ID, bind & return Adventure model.
func (c *Collection) GetOne(_ context.Context, id string) (*model.Adventure, error) {
	var adventure model.Adventure
	adventure.Adventure_id = id
	adventure.User_id = "652edaa67a75034ea37c6652"
	adventure.Activity = "Test"
	return &adventure, nil
}

func (c *Collection) GetAll(_ context.Context, user_id string) ([]*model.Adventure, error) {
	var adventures []*model.Adventure
	return adventures, nil
}

func (c *Collection) Delete(_ context.Context, id string) error {
	return nil
}

func (c *Collection) Update(_ context.Context, adventure *model.Adventure) error {
	return nil
}

func (c *Collection) Create(_ context.Context, adventure *model.Adventure) error {
	adventure.Adventure_id = "656001daf827a04b7a66bafa"
	return nil
}
