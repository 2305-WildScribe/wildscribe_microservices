package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Rating defines an individual rating created by a user.
type User struct {
	User_id  primitive.ObjectID `json:"user_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" validate:"required"`
	Email    string             `json:"email" validate:"required"`
	Password string             `json:"password" validate:"required"`
}
