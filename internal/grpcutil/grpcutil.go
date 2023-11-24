package grpcutil

import (
	"context"
	// "math/rand"
	// "yourpackage/discovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ServiceConnection connects to a fixed URL for the specified service.
func ServiceConnection(ctx context.Context, serviceName string) (*grpc.ClientConn, error) {
	// Specify the fixed URL for the service

	return grpc.Dial(serviceName, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
