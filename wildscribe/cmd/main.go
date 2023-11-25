package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"wildscribe.com/wildscribe/internal/controller"
	grpchandler "wildscribe.com/wildscribe/internal/handler/grpc"

	"github.com/gin-gonic/gin"
	"wildscribe.com/gen"
	adventuregateway "wildscribe.com/wildscribe/internal/gateway/adventure/grpc"
	usergateway "wildscribe.com/wildscribe/internal/gateway/user/grpc"
	"wildscribe.com/wildscribe/internal/handler/gin_handler"
	"wildscribe.com/wildscribe/internal/routes"

)

func main() {
	var port string
	var address string
	log.Println("Starting main Wildscribe service...")
	env := os.Getenv("ENV")

	if env == "PROD" {
		port = os.Getenv("PORT")
		address = "0.0.0.0"
	} else {
		port = "8080"
		address = "0.0.0.0"
	}

	// Adventure grpc gateway setup
	log.Println("Setting up Adventure Gateway")
	adventureGateway := adventuregateway.NewAdventureGateway("0.0.0.0:8083")
	log.Println("Done!")
	// User grpc gateway setup
	log.Println("Setting up User Gateway")
	userGateway := usergateway.NewUserGateway("https://wildscribe-user-service-97db90e759bf.herokuapp.com")
	log.Println("Done!")
	// Setup controller for main func
	log.Println("Setting controller")
	ctrl := controller.New(adventureGateway, userGateway)
	log.Println("Done!")

	// Setup Gin router for main controller
	route := fmt.Sprintf("%s:%s", address, port)

	log.Printf("Environment: %s, Address: %s, Port: %s, Route: %s\n", env, address, port, route)

	router := gin.Default()

	log.Println("Setting handler")
	handler := gin_handler.NewGinHandler(ctrl)
	log.Println("Done!")

	log.Println("Setting routes")
	routes.Routes(router, handler)
	log.Println("Done!")
	server := &http.Server{
		Addr:    route,
		Handler: router,
	}
	go func() {
		log.Println("Starting HTTP service")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server startup failed: %v", err)
		}
	}()

	srv := grpc.NewServer()
	go func() {
		log.Println("Setting up gRPC handler")
		ghandler := grpchandler.New(ctrl)
		log.Println("Done!")

		log.Println("Setting up gRPC listen address")
		lis, err := net.Listen("tcp", "0.0.0.0:8084")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Println("Done!")

		log.Println("Setting up gRPC server")
		log.Println("Done!")

		log.Println("Starting up gRPC server")
		gen.RegisterAdventureServiceServer(srv, ghandler)
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("gRPC server failed to serve: %v", err)
		}
		log.Println("gRPC server stopped")
	}()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	// Gracefully shut down the HTTP server
	log.Println("Stopping HTTP service")
	if err := server.Close(); err != nil {
		log.Fatalf("HTTP server shutdown failed: %v", err)
	}
	log.Println("HTTP service gracefully stopped")

	// Gracefully shut down the gRPC server
	log.Println("Stopping gRPC server")
	srv.GracefulStop()
	log.Println("gRPC server gracefully stopped")
}
