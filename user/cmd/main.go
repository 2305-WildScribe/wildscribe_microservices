package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "gopkg.in/yaml.v3"
	"log"
	"os"
	"wildscribe.com/user/internal/controller/user"
	"wildscribe.com/user/internal/handler/gin_handler"
	"wildscribe.com/user/internal/repository/mongoDB"
	"wildscribe.com/user/internal/routes"
)

func main() {
	var port string
	var address string
	log.Println("Starting wildscribe adventure service...")
	env := os.Getenv("ENV")

	if env == "PROD" {
		port = os.Getenv("PORT")
		address = "0.0.0.0"
	} else {
		port = "8082"
		address = "0.0.0.0"
	}

	route := fmt.Sprintf("%s:%s", address, port)
	log.Printf("Environment: %s, Address: %s, Port: %s, Route: %s\n", env, address, port, route)
	router := gin.Default()
	log.Println("Connecting to MongoDB")
	db := mongoDB.ConnectDB()
	log.Println("Done!")

	log.Println("Setting up collection")
	repo := mongoDB.NewCollection(db)
	log.Println("Done!")

	log.Println("Setting controller")
	ctrl := user.New(repo)
	log.Println("Done!")

	log.Println("Setting handler")
	handler := gin_handler.New(ctrl)
	log.Println("Done!")

	log.Println("Setting routes")
	routes.UserRoute(router, handler)
	log.Println("Done!")

	log.Println("Starting service")
	router.Run(route)
	log.Println("Done! Service is live!")
}
