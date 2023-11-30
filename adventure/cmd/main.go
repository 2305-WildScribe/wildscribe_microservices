package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"

	"wildscribe.com/adventure/internal/controller"
	grpchandler "wildscribe.com/adventure/internal/handler/grpc"
	database "wildscribe.com/adventure/internal/repository/mongoDB"
	"wildscribe.com/gen"
)

func main() {
	var port string
	var address string
	log.Println("Starting wildscribe adventure service...")
	env := os.Getenv("ENV")

	if env == "PROD" {
		port = os.Getenv("PORT")
		address = os.Getenv("ADDRESS")
	} else {
		port = "8083"
		address = "0.0.0.0"
	}

	route := fmt.Sprintf("%s:%s", address, port)
	log.Printf("Environment: %s, Address: %s, Port: %s, Route: %s\n", env, address, port, route)

	log.Println("Connecting to MongoDB")
	db := database.ConnectDB()
	log.Println("Done!")

	log.Println("Setting up collection")
	repo := database.NewCollection(db)
	log.Println("Done!")

	log.Println("Setting controller")
	svc := controller.New(repo)
	log.Println("Done!")

	log.Println("Setting Handler")
	h := grpchandler.New(svc)
	log.Println("Done!")

	log.Println("Setting up route")
	lis, err := net.Listen("tcp", route)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Done!")

	log.Println("Setting service")
	srv := grpc.NewServer()
	log.Println("Done!")

	log.Println("Starting service")
	gen.RegisterAdventureServiceServer(srv, h)
	srv.Serve(lis)

	log.Println("Service Shutting down! :(")

}
