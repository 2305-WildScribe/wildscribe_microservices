package main_test

import (
	"context"

	// "github.com/stretchr/testify/assert"
	"log"
	"net"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"wildscribe.com/adventure/internal/controller"
	grpchandler "wildscribe.com/adventure/internal/handler/grpc"
	database "wildscribe.com/adventure/internal/repository/mockDB"
	"wildscribe.com/gen"
)

// ... (your imports and other code)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var adventureClient gen.AdventureServiceClient

func TestMain(m *testing.M) {
	ctx := context.Background()

	adventureSrv := startAdventureService(ctx)
	defer adventureSrv.GracefulStop()
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	adventureConn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), opts)
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer adventureConn.Close()

	adventureClient = gen.NewAdventureServiceClient(adventureConn)

	os.Exit(m.Run())
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func startAdventureService(ctx context.Context) *grpc.Server {
	db := database.MockCollection()

	log.Println("Setting controller")
	svc := controller.New(db)

	log.Println("Done!")

	log.Println("Setting Handler")
	h := grpchandler.New(svc)
	log.Println("Done!")

	log.Println("Setting service")
	lis = bufconn.Listen(bufSize)
	srv := grpc.NewServer()
	log.Println("Done!")

	log.Println("Starting service")
	gen.RegisterAdventureServiceServer(srv, h)
	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Printf("Server stopped: %v", err)
		}
	}()
	return srv
}

func TestGetAdventure(t *testing.T) {
	log.Println("Testing Get Adventure")
	ctx := context.Background()

	adventure := &gen.Adventure{
		AdventureId: "652edaa67a75034ea37c6652",
		UserId:      "652edaa67a75034ea37c6679",
		Activity:    "Test",
	}

	getAdventureResp, err := adventureClient.GetAdventure(ctx, &gen.GetAdventureRequest{AdventureId: adventure.AdventureId})
	if err != nil {
		log.Fatalf("Get Adventure: %v", err)
	}
	if diff := cmp.Diff(getAdventureResp.Adventure, adventure, cmpopts.IgnoreUnexported(gen.Adventure{})); diff != "" {
		log.Fatalf("Get adventure mismatch: %s", diff)
	}
}
