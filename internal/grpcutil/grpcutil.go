package grpcutil

import (
	"context"
	"fmt"
	"log"

	// "math/rand"
	// "yourpackage/discovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ServiceConnection connects to a fixed URL for the specified service.
func ServiceConnection(ctx context.Context, serviceAddress string) (*grpc.ClientConn, error) {
	// Specify the fixed URL for the service
	log.Printf("GrpcUtil::ServiceConnection: Connected to %s", serviceAddress)
	client, err := grpc.Dial(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println(fmt.Errorf("GrpcUtil::ServiceConnection: Error connecting to service: %w", err))
	log.Println(client)
	return client, err
}
