package main

import (
	// External Dependencies
	"fmt"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"os"

	// Internal Dependencies
	"wildscribe.com/adventure/internal/controller"
	grpchandler "wildscribe.com/adventure/internal/handler/grpc"
	database "wildscribe.com/adventure/internal/repository/mongoDB"
	"wildscribe.com/gen"
)

func main() {
	var cfg ServiceConfig

	log.Println("Starting wildscribe adventure service...")

	env := os.Getenv("ENV")
	if env == "PROD" {

		cfg.APIConfig.Port = os.Getenv("PORT")
		cfg.APIConfig.Address = os.Getenv("ADDRESS")
	} else {
		config, err := os.Open("configs/base.yaml")
		if err != nil {
			panic(err)
		}

		defer config.Close()
		if err := yaml.NewDecoder(config).Decode(&cfg); err != nil {
			panic(err)
		}
	}

	route := fmt.Sprintf("%s:%s", cfg.APIConfig.Address, cfg.APIConfig.Port)
	log.Printf("Environment: %s, Route: %s\n", env, route)

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
