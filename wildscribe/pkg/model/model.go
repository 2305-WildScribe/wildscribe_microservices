package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User Model with UserID, and Name.
type User struct {
	UserID   primitive.ObjectID `json:"user_id"`
	Name     string             `json:"name,omitempty"`
	Email    string             `json:"email,omitempty"`
	Password string             `json:"password,omitempty"`
}
