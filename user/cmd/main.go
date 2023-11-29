package main

import (
	"fmt"
	"net"

	"log"
	"os"

	"google.golang.org/grpc"
	"wildscribe.com/gen"
	"wildscribe.com/user/internal/controller"
	grpchandler "wildscribe.com/user/internal/handler/grpc"
	database "wildscribe.com/user/internal/repository/mongoDB"
)

func main() {
	var port string
	var address string
	log.Println("Starting wildscribe user service...")
	env := os.Getenv("ENV")

	if env == "PROD" {
		port = os.Getenv("PORT")
		address = os.Getenv("ADDRESS")
	} else {
		port = "8082"
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
	ctrl := controller.New(repo)
	log.Println("Done!")

	log.Println("Setting handler")
	handler := grpchandler.New(ctrl)
	log.Println("Done!")

	log.Println("Setting up gRPC server")
	lis, err := net.Listen("tcp", route)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Done!")

	srv := grpc.NewServer()
	gen.RegisterUserServiceServer(srv, handler)
	srv.Serve(lis)
}
