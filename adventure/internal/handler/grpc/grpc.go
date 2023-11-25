package grpc

import (
	"context"
	"fmt"
	"log"

	"wildscribe.com/adventure/internal/controller"
	"wildscribe.com/adventure/pkg/model"
	"wildscribe.com/gen"
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
	m, err := h.svc.Show(ctx, req.AdventureId)
	if err != nil {
		new_error := fmt.Errorf("grpc_Handler::GetAdventure: Failed to fetch adventure: %w", err)
		log.Println(new_error)
		return nil, new_error
	}
	return &gen.GetAdventureResponse{Adventure: model.AdventureToProto(m)}, nil
}

func (h *Handler) GetAllAdventures(ctx context.Context, req *gen.GetAllAdventuresRequest) (*gen.GetAllAdventuresResponse, error) {
	adventures, err := h.svc.Index(ctx, req.UserId)
	if err != nil {
		new_error := fmt.Errorf("grpc_Handler::GetAllAdventures: Failed to fetch adventures: %w", err)
		log.Println(new_error)
		return nil, new_error
	}
	protoAdventures := model.AdventureSliceToProto(adventures)
	return &gen.GetAllAdventuresResponse{Adventures: protoAdventures}, nil
}

func (h *Handler) CreateAdventure(ctx context.Context, req *gen.CreateAdventureRequest) (*gen.CreateAdventureResponse, error) {
	adventure, err := h.svc.Create(ctx, model.AdventureFromProto(req.Adventure))
	if err != nil {
		new_error := fmt.Errorf("grpc_Handler::CreateAdventures: Failed to create adventure: %w", err)
		log.Println(new_error)
		return nil, new_error
	}
	return &gen.CreateAdventureResponse{Adventure: model.AdventureToProto(adventure)}, nil
}

func (h *Handler) UpdateAdventure(ctx context.Context, req *gen.UpdateAdventureRequest) (*gen.UpdateAdventureResponse, error) {
	adventure, err := h.svc.Update(ctx, model.AdventureFromProto(req.Adventure))
	if err != nil {
		new_error := fmt.Errorf("grpc_Handler::UpdateAdventures: Failed to Update adventure: %w", err)
		log.Println(new_error)
		return nil, new_error
	}
	return &gen.UpdateAdventureResponse{Adventure: model.AdventureToProto(adventure)}, nil
}

func (h *Handler) DeleteAdventure(ctx context.Context, req *gen.DeleteAdventureRequest) (*gen.DeleteAdventureResponse, error) {
	err := h.svc.Delete(ctx, req.AdventureId)
	if err != nil {
		new_error := fmt.Errorf("grpc_Handler::DeleteAdventures: Failed to Delete adventure: %w", err)
		log.Println(new_error)
		return nil, new_error
	}
	return &gen.DeleteAdventureResponse{AdventureId: req.AdventureId}, nil
}
