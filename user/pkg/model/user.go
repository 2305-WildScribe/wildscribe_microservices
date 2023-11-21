package model

// Rating defines an individual rating created by a user.
type User struct {
	User_id  string `bson:"_id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
