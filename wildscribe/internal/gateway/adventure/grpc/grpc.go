package grpc

import (
	"context"
	"fmt"
	"log"
	// "google.golang.org/grpc"
	// "go.mongodb.org/mongo-driver/mongo/address"
	"wildscribe.com/adventure/pkg/model"
	"wildscribe.com/gen"
	"wildscribe.com/internal/grpcutil"
	// "wildscribe.com/pkg/discovery"
)

type Gateway struct {
	address         string
	adventureClient gen.AdventureServiceClient
}

// New creates a new gRPC gateway for a movie metadata
// service.
func NewAdventureGateway(addr string) *Gateway {
	conn, err := grpcutil.ServiceConnection(context.Background(), "0.0.0.0:8083")
	if err != nil {
		log.Fatalf("Failed to connect to Adventure service: %v", err)
	}
	return &Gateway{addr, gen.NewAdventureServiceClient(conn)}
}

// Get returns movie metadata by a movie id.
func (g *Gateway) GetAdventure(ctx context.Context, adventure_id string) (*model.Adventure, error) {
	resp, err := g.adventureClient.GetAdventure(ctx, &gen.GetAdventureRequest{AdventureId: adventure_id})
	if err != nil {
		new_error := fmt.Errorf("AdvGrpcGateway::GetAdventure: Error getting adventure: %w", err)
		return nil, new_error
	}
	return model.AdventureFromProto(resp.Adventure), nil
}

func (g *Gateway) GetAllAdventures(ctx context.Context, user_id string) ([]*model.Adventure, error) {
	resp, err := g.adventureClient.GetAllAdventures(ctx, &gen.GetAllAdventuresRequest{UserId: user_id})
	if err != nil {
		new_error := fmt.Errorf("AdvGrpcGateway::GetAllAdventures: Error getting adventures: %w", err)
		return nil, new_error
	}
	adventures := model.AdventureSliceFromProto(resp.Adventures)
	log.Println(adventures)

	return adventures, nil
}

func (g *Gateway) CreateAdventure(ctx context.Context, adventure *model.Adventure) (*model.Adventure, error) {
	resp, err := g.adventureClient.CreateAdventure(ctx, &gen.CreateAdventureRequest{Adventure: model.AdventureToProto(adventure)})
	if err != nil {
		new_error := fmt.Errorf("AdvGrpcGateway::CreateAdventure: Error creating adventure: %w", err)
		return nil, new_error
	}
	return model.AdventureFromProto(resp.Adventure), nil
}

func (g *Gateway) UpdateAdventure(ctx context.Context, adventure *model.Adventure) (*model.Adventure, error) {
	resp, err := g.adventureClient.UpdateAdventure(ctx, &gen.UpdateAdventureRequest{Adventure: model.AdventureToProto(adventure)})
	if err != nil {
		new_error := fmt.Errorf("AdvGrpcGateway::UpdateAdventure: Error creating adventure: %w", err)
		return nil, new_error
	}
	return model.AdventureFromProto(resp.Adventure), nil
}

func (g *Gateway) DeleteAdventure(ctx context.Context, adventure_id string) (string, error) {
	resp, err := g.adventureClient.DeleteAdventure(ctx, &gen.DeleteAdventureRequest{AdventureId: adventure_id})
	if err != nil {
		new_error := fmt.Errorf("AdvGrpcGateway::DeleteAdventure: Error deleting adventure: %w", err)
		return adventure_id, new_error
	}
	return resp.AdventureId, nil
}
