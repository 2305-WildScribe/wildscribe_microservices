package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"

	"wildscribe.com/wildscribe/internal/controller"
	adventuregateway "wildscribe.com/wildscribe/internal/gateway/adventure/grpc"
	usergateway "wildscribe.com/wildscribe/internal/gateway/user/grpc"
	"wildscribe.com/wildscribe/internal/handler/gin_handler"
	"wildscribe.com/wildscribe/internal/routes"
)

func main() {
	var port string
	var address string
	var adventureAddress string
	var userAddress string
	// var grpcListen string
	log.Println("Starting main Wildscribe service...")
	env := os.Getenv("ENV")

	if env == "PROD" {
		port = os.Getenv("PORT")
		address = os.Getenv("ADDRESS")
		adventureAddress = os.Getenv("ADVENTUREGATEWAY")
		userAddress = os.Getenv("USERGATEWAY")
	} else {
		port = "8080"
		address = "0.0.0.0"
		adventureAddress = "adventure-microservice:8083"
		userAddress = "0.0.0.0:8082"
	}

	// Adventure grpc gateway setup
	log.Println("Setting up Adventure Gateway")
	adventureGateway := adventuregateway.NewAdventureGateway(adventureAddress)
	log.Println("Done!")
	// User grpc gateway setup
	log.Println("Setting up User Gateway")
	userGateway := usergateway.NewUserGateway(userAddress)
	log.Println("Done!")
	// Setup controller for main func
	log.Println("Setting controller")
	ctrl := controller.New(adventureGateway, userGateway)
	log.Println("Done!")

	// Setup Gin router for main controller
	route := fmt.Sprintf("%s:%s", address, port)

	log.Printf("Environment: %s, Address: %s, Port: %s, Route: %s\n", env, address, port, route)

	router := gin.Default()
	router.Use(cors.Default())
	router.UseH2C = true
	log.Println("Setting handler")
	handler := gin_handler.NewGinHandler(ctrl)
	log.Println("Done!")

	log.Println("Setting routes")
	routes.Routes(router, handler)
	log.Println("Done!")

	log.Println("Starting Service!")
	router.Run()
}
