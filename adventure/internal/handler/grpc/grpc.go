package grpc

import (
	// External Dependencies
	"context"
	"fmt"
	//
	"wildscribe.com/adventure/internal/controller"
	"wildscribe.com/adventure/pkg/model"
	"wildscribe.com/gen"
)

// Handler defines a adventure gRPC handler.
type Handler struct {
	gen.UnimplementedAdventureServiceServer
	svc *controller.Controller
}

// New creates a new adventure gRPC handler.
func New(ctrl *controller.Controller) *Handler {
	return &Handler{svc: ctrl}
}

// GetAdventure returns single adventure by adventureId
func (h *Handler) GetAdventure(ctx context.Context, req *gen.GetAdventureRequest) (*gen.GetAdventureResponse, error) {
	m, err := h.svc.Show(ctx, req.AdventureId)
	if err != nil {
		return nil, fmt.Errorf("grpcHandler::GetAdventure: Failed to fetch adventure: %w", err)
	}
	return &gen.GetAdventureResponse{Adventure: model.AdventureToProto(m)}, nil
}

// GetAllAdventures Takes userId and returns a slice of Adventure models for that userId
func (h *Handler) GetAllAdventures(ctx context.Context, req *gen.GetAllAdventuresRequest) (*gen.GetAllAdventuresResponse, error) {
	adventures, err := h.svc.Index(ctx, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("grpcHandler::GetAllAdventures: Failed to fetch adventures: %w", err)
	}
	protoAdventures := model.AdventureSliceToProto(adventures)
	return &gen.GetAllAdventuresResponse{Adventures: protoAdventures}, nil
}

func (h *Handler) CreateAdventure(ctx context.Context, req *gen.CreateAdventureRequest) (*gen.CreateAdventureResponse, error) {
	adventure, err := h.svc.Create(ctx, model.AdventureFromProto(req.Adventure))
	if err != nil {
		return nil, fmt.Errorf("grpcHandler::CreateAdventures: Failed to create adventure: %w", err)
	}
	return &gen.CreateAdventureResponse{Adventure: model.AdventureToProto(adventure)}, nil
}

func (h *Handler) UpdateAdventure(ctx context.Context, req *gen.UpdateAdventureRequest) (*gen.UpdateAdventureResponse, error) {
	adventure, err := h.svc.Update(ctx, model.AdventureFromProto(req.Adventure))
	if err != nil {
		return nil, fmt.Errorf("grpcHandler::UpdateAdventures: Failed to Update adventure: %w", err)
	}
	return &gen.UpdateAdventureResponse{Adventure: model.AdventureToProto(adventure)}, nil
}

func (h *Handler) DeleteAdventure(ctx context.Context, req *gen.DeleteAdventureRequest) (*gen.DeleteAdventureResponse, error) {
	adventure_id, err := h.svc.Delete(ctx, req.AdventureId)
	if err != nil {
		return nil, fmt.Errorf("grpcHandler::DeleteAdventures: Failed to Delete adventure: %w", err)
	}
	return &gen.DeleteAdventureResponse{AdventureId: adventure_id}, nil
}
