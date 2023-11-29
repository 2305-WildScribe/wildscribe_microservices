package grpc

import (
	"context"
	"fmt"
	"log"
	"wildscribe.com/gen"
	"wildscribe.com/internal/grpcutil"
	"wildscribe.com/user/pkg/model"
)

type Gateway struct {
	address    string
	userClient gen.UserServiceClient
}

// New creates a new gRPC gateway for a user service.
func NewUserGateway(addr string) *Gateway {
	conn, err := grpcutil.ServiceConnection(context.Background(), addr)
	if err != nil {
		log.Printf("Failed to connect to User service: %v", err)
	}
	return &Gateway{addr, gen.NewUserServiceClient(conn)}
}

// Get returns movie metadata by a movie id.
func (g *Gateway) LoginUser(ctx context.Context, user *model.User) (*model.User, error) {
	resp, err := g.userClient.LoginUser(ctx, &gen.LoginUserRequest{Email: user.Email, Password: user.Password})
	if err != nil {
		new_error := fmt.Errorf("UserGrpcGateway::LoginUser: Error getting User: %w", err)
		return nil, new_error
	}
	resp_user := model.UserFromProto(resp.User)
	log.Println(resp_user)
	resp_user.Email = ""
	resp_user.Password = ""
	return resp_user, nil
}

func (g *Gateway) ValidateUser(ctx context.Context, user_id string) (string, error) {
	resp, err := g.userClient.ValidateUser(ctx, &gen.ValidateUserRequest{UserId: user_id})
	if err != nil {
		new_error := fmt.Errorf("UserGrpcGateway::ValidateUser: Error getting User: %w", err)
		return "", new_error
	}
	return resp.UserId, nil
}

func (g *Gateway) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	resp, err := g.userClient.CreateUser(ctx, &gen.CreateUserRequest{User: model.UserToProto(user)})
	if err != nil {
		new_error := fmt.Errorf("UserGrpcGateway::CreateUser: Error getting User: %w", err)
		return nil, new_error
	}
	return model.UserFromProto(resp.User), nil
}

func (g *Gateway) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	resp, err := g.userClient.UpdateUser(ctx, &gen.UpdateUserRequest{User: model.UserToProto(user)})
	if err != nil {
		new_error := fmt.Errorf("UserGrpcGateway::UpdateUser: Error getting User: %w", err)
		return nil, new_error
	}
	return model.UserFromProto(resp.User), nil
}

func (g *Gateway) DeleteUser(ctx context.Context, user_id string) (string, error) {
	resp, err := g.userClient.DeleteUser(ctx, &gen.DeleteUserRequest{UserId: user_id})
	if err != nil {
		new_error := fmt.Errorf("UserGrpcGateway::DeleteUser: Error getting User: %w", err)
		return "", new_error
	}
	return resp.UserId, nil
}
