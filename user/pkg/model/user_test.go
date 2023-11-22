package model

import (
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestNewUser(t *testing.T) {
	// Create a new User
	user := NewUser("name", "email", "password")

	// Verify the User fields
	if user.Name != "name" {
		t.Errorf("Expected Name to be 'name', but got %s", user.Name)
	}
	if user.Email != "email" {
		t.Errorf("Expected Email to be 'email', but got %s", user.Email)
	}
	if user.Password != "password" {
		t.Errorf("Expected Password to be 'password', but got %s", user.Password)
	}
}

func TestUserFields(t *testing.T) {
	user := NewUser("name", "email", "password")

	if user.Name != "name" {
		t.Errorf("Expected User Name to be name, but got %v", user.Name)
	}
	if user.Email != "email" {
		t.Errorf("Expected User Email to be email, but got %v", user.Email)
	}
	if user.Password != "password" {
		t.Errorf("Expected User Password to be password, but got %v", user.Password)
	}
}
