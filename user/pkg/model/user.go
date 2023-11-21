package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Rating defines an individual rating created by a user.
type User struct {
	User_id  primitive.ObjectID `json:"user_id,omitempty" json:"_id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}
