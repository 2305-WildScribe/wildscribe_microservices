package grpc

import (
	"context"
	"fmt"
	"log"
	"wildscribe.com/gen"
	"wildscribe.com/user/internal/controller"
	"wildscribe.com/user/pkg/model"
)

type Handler struct {
	gen.UnimplementedUserServiceServer
	svc *controller.Controller
}

func New(ctrl *controller.Controller) *Handler {
	return &Handler{svc: ctrl}
}

func (h *Handler) LoginUser(ctx context.Context, req *gen.LoginUserRequest) (*gen.LoginUserResponse, error) {
	user, err := h.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		new_error := fmt.Errorf("GrpcHandler::LoginUser: failed to fetch user %w", err)
		log.Println(new_error)
		return nil, new_error
	}
	return &gen.LoginUserResponse{User: model.UserToProto(user)}, nil
}

func (h *Handler) ValidateUser(ctx context.Context, req *gen.ValidateUserRequest) (*gen.ValidateUserResponse, error) {
	userBool, err := h.svc.Validate(ctx, req.UserId)
	if err != nil {
		new_error := fmt.Errorf("GrpcHandler::ValidateUser: failed to validate user %w", err)
		log.Println(new_error)
		return nil, new_error
	}
	if !userBool {
		new_error := fmt.Errorf("GrpcHandler::ValidateUser: Invalid User ID %w", err)
		log.Println(new_error)
		return nil, new_error
	}
	return &gen.ValidateUserResponse{UserId: req.UserId}, nil
}

func (h *Handler) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	user, err := h.svc.Create(ctx, model.UserFromProto(req.User))
	if err != nil {
		new_error := fmt.Errorf("GrpcHandler::CreateUser: failed to create user %w", err)
		log.Println(new_error)
		return nil, new_error
	}
	return &gen.CreateUserResponse{User: model.UserToProto(user)}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *gen.UpdateUserRequest) (*gen.UpdateUserResponse, error) {
	user, err := h.svc.Update(ctx, model.UserFromProto(req.User))
	if err != nil {
		new_error := fmt.Errorf("GrpcHandler::UpdateUser: failed to update user %w", err)
		log.Println(new_error)
		return nil, new_error
	}
	return &gen.UpdateUserResponse{User: model.UserToProto(user)}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *gen.DeleteUserRequest) (*gen.DeleteUserResponse, error) {
	err := h.svc.Delete(ctx, req.UserId)
	if err != nil {
		new_error := fmt.Errorf("GrpcHandler::DeleteUser: failed to delete user %w", err)
		log.Println(new_error)
		return nil, new_error
	}
	return &gen.DeleteUserResponse{UserId: req.UserId}, nil
}
