package grpcutil

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ServiceConnection connects to a fixed URL for the specified service.
func ServiceConnection(ctx context.Context, host string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if host != "" {
		opts = append(opts, grpc.WithAuthority(host))
	}
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if insecure {
	// 	opts = append(opts, grpc.WithInsecure())
	// }
	// Specify the fixed URL for the service
	client, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Println(fmt.Errorf("GrpcUtil::ServiceConnection: Error connecting to service: %w", err))
	}
	log.Printf("GrpcUtil::ServiceConnection: Connected to %s", host)
	return client, err
}
