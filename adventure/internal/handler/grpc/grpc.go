package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// "wildscribe.com/adventure/internal/controller/adventure"
	"wildscribe.com/adventure/internal/controller/adventure"
	// "wildscribe.com/adventure/internal/repository"
	"wildscribe.com/adventure/pkg/model"
	"wildscribe.com/gen"
)

// Handler defines a movie metadata gRPC handler.
type Handler struct {
	gen.UnimplementedAdventureServiceServer
	svc *adventure.Controller
}

// New creates a new movie metadata gRPC handler.
func New(ctrl *adventure.Controller) *Handler {
	return &Handler{svc: ctrl}
}

// GetAdventureByID returns movie Adventure by id.
func (h *Handler) GetAdventure(ctx context.Context, req *gen.GetAdventureRequest) (*gen.GetAdventureResponse, error) {
	if req == nil || req.AdventureId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.svc.Show(ctx, req.AdventureId)
	if err != nil && errors.Is(err, adventure.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetAdventureResponse{Adventure: model.AdventureToProto(m)}, nil
}
