package model

import (
	"wildscribe.com/gen"
)

func UserToProto(user *User) *gen.User {
	return &gen.User{
		UserId:   user.User_id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func UserFromProto(user *gen.User) *User {
	return &User{
		User_id:  user.UserId,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}