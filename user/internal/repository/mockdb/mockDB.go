package mockdb

import (
	"context"
	"errors"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"wildscribe.com/user/internal/middleware/bcrypt_middleware"
	"wildscribe.com/user/pkg/model"
)

// // Set Database stuct
//
//	type Database struct {
//		db *mongo.Client
//	}
var ErrNotFound = errors.New("users not found for a record")

// Set Collection struct
type Collection struct {
	collection string
}

// Sets the collection to "metadata"
func MockCollection() *Collection {
	return &Collection{
		collection: "Test",
	}
}

// Set user model

// Get a single collection from the ID, bind & return user model.
func (c *Collection) Get(_ context.Context, email string) (*model.User, error) {
	var user model.User
	user.Name = "Ian"
	user.User_id = "65330eb5fcb829e722f7c40c"
	user.Email = "me@gmail.com"
	user.Password = bcrypt_middleware.HashAndSalt([]byte("password"))
	return &user, nil
}
