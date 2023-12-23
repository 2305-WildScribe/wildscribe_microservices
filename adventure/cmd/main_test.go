package main_test

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"log"
	"net"
	"os"
	"testing"

	"wildscribe.com/adventure/internal/controller"
	grpchandler "wildscribe.com/adventure/internal/handler/grpc"
	database "wildscribe.com/adventure/internal/repository/mockDB"
	"wildscribe.com/gen"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener
var adventureClient gen.AdventureServiceClient
var ctx = context.Background()

var adventureId string

func TestMain(m *testing.M) {

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
	db := database.SetupMockDB("adventures")

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

func TestCreateAdventure(t *testing.T) {
	log.Println("Testing Create Adventure")
	expectedResponse := &gen.Adventure{
		UserId:   "652edaa67a75034ea37c6679",
		Activity: "Test",
	}
	adventureRequest := &gen.Adventure{
		UserId:   "652edaa67a75034ea37c6679",
		Activity: "Test",
	}

	createAdventureResp, err := adventureClient.CreateAdventure(ctx, &gen.CreateAdventureRequest{Adventure: adventureRequest})
	if err != nil {
		log.Fatalf("Get Adventure: %v", err)
	}

	if diff := cmp.Diff(createAdventureResp.Adventure.UserId, expectedResponse.UserId, cmpopts.IgnoreUnexported(gen.Adventure{})); diff != "" {
		log.Fatalf("Get adventure mismatch: %s", diff)
	}

	adventureId = createAdventureResp.Adventure.AdventureId
}

func TestGetAdventure(t *testing.T) {
	log.Println("Testing Get Adventure")

	adventure := &gen.Adventure{
		AdventureId: adventureId,
		UserId:      "652edaa67a75034ea37c6679",
		Activity:    "Test",
	}

	adventureResponse, err := adventureClient.GetAdventure(ctx, &gen.GetAdventureRequest{AdventureId: adventure.AdventureId})
	if err != nil {
		log.Fatalf("Get Adventure: %v", err)
	}
	if diff := cmp.Diff(adventureResponse.Adventure, adventure, cmpopts.IgnoreUnexported(gen.Adventure{})); diff != "" {
		log.Fatalf("Get adventure mismatch: %s", diff)
	}
}

func TestGetAllAdventures(t *testing.T) {
	log.Println("Testing Get All Adventures")

	adventure := &gen.Adventure{
		AdventureId: adventureId,
		UserId:      "652edaa67a75034ea37c6679",
		Activity:    "Test",
	}

	getAdventureResp, err := adventureClient.GetAllAdventures(ctx, &gen.GetAllAdventuresRequest{UserId: adventure.UserId})
	if err != nil {
		log.Fatalf("Get All Adventure: %v", err)
	}
	if diff := cmp.Diff(getAdventureResp.Adventures[0], adventure, cmpopts.IgnoreUnexported(gen.Adventure{})); diff != "" {
		log.Fatalf("Get All adventure mismatch: %s", diff)
	}
}

func TestUpdateAdventure(t *testing.T) {

	expectedResponse := &gen.Adventure{
		AdventureId: adventureId,
		UserId:      "652edaa67a75034ea37c6679",
		Activity:    "Updated Test",
	}

	log.Println("Testing Update Adventure")
	updateAdventureRequest := &gen.Adventure{
		AdventureId: adventureId,
		UserId:      "652edaa67a75034ea37c6679",
		Activity:    "Updated Test",
	}

	updateAdventureResponse, err := adventureClient.UpdateAdventure(ctx, &gen.UpdateAdventureRequest{Adventure: updateAdventureRequest})
	if err != nil {
		log.Fatalf("Update Adventure: %v", err)
	}
	if diff := cmp.Diff(updateAdventureResponse.Adventure, expectedResponse, cmpopts.IgnoreUnexported(gen.Adventure{})); diff != "" {
		log.Fatalf("Update adventure mismatch: %s", diff)
	}
}

func TestDeleteAdventure(t *testing.T) {
	log.Println("Testing Delete Adventure")

	expectedResponse := &gen.Adventure{
		AdventureId: adventureId,
		UserId:      "652edaa67a75034ea37c6679",
		Activity:    "Test",
	}

	deleteAdventureResponse, err := adventureClient.DeleteAdventure(ctx, &gen.DeleteAdventureRequest{AdventureId: adventureId})
	if err != nil {
		log.Fatalf("Delete Adventure: %v", err)
	}
	if diff := cmp.Diff(deleteAdventureResponse.AdventureId, expectedResponse.AdventureId, cmpopts.IgnoreUnexported(gen.Adventure{})); diff != "" {
		log.Fatalf("Delete adventure mismatch: %s", diff)
	}
}
