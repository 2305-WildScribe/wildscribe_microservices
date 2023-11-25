package grpc

import (
	"context"
	"fmt"
	"log"

	"wildscribe.com/adventure/pkg/model"
	"wildscribe.com/gen"
	"wildscribe.com/wildscribe/internal/controller"
)

// Handler defines a movie metadata gRPC handler.
type Handler struct {
	gen.UnimplementedAdventureServiceServer
	svc *controller.Controller
}

// New creates a new movie metadata gRPC handler.
func New(ctrl *controller.Controller) *Handler {
	return &Handler{svc: ctrl}
}

// GetAdventureByID returns movie Adventure by id.
func (h *Handler) GetAdventure(ctx context.Context, req *gen.GetAdventureRequest) (*gen.GetAdventureResponse, error) {
	m, err := h.svc.GetAdventure(ctx, req.AdventureId)
	if err != nil {
		new_error := fmt.Errorf("grpc_Handler::GetAdventure: Failed to fetch Adventure: %w", err)
		log.Println(new_error)
		return nil, new_error
	}
	return &gen.GetAdventureResponse{Adventure: model.AdventureToProto(m)}, nil
}

func (h *Handler) GetAllAdventures(ctx context.Context, req *gen.GetAllAdventuresRequest) (*gen.GetAllAdventuresResponse, error) {
	adventures, err := h.svc.GetAllAdventures(ctx, req.UserId)
	if err != nil {
		new_error := fmt.Errorf("grpc_Handler::GetAllAdventures: Failed to fetch Adventure: %w", err)
		log.Println(new_error)
		return nil, new_error
	}
	protoAdventures := model.AdventureSliceToProto(adventures)
	return &gen.GetAllAdventuresResponse{Adventures: protoAdventures}, err
}
