package model

// import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	User_id  string `json:"user_id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
