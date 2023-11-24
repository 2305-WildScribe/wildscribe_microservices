package grpc

import (
	"context"
	// "google.golang.org/grpc"
	"wildscribe.com/adventure/pkg/model"
	"wildscribe.com/gen"
	"wildscribe.com/internal/grpcutil"
	// "wildscribe.com/pkg/discovery"
)

type Gateway struct {
	// registry discovery.Registry
}

// New creates a new gRPC gateway for a movie metadata
// service.
func NewAdvGrpcGateway() *Gateway {
	return &Gateway{}
}

// Get returns movie metadata by a movie id.
func (g *Gateway) GetAdventure(ctx context.Context, adventure_id string) (*model.Adventure, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "Adventure")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := gen.NewAdventureServiceClient(conn)
	resp, err := client.GetAdventure(ctx, &gen.GetAdventureRequest{AdventureId: adventure_id})
	if err != nil {
		return nil, err
	}
	return model.AdventureFromProto(resp.Adventure), nil
}
