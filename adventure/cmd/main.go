package main

import (
	"fmt"
	// "github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"wildscribe.com/adventure/internal/controller/adventure"
	"wildscribe.com/gen"
	// "wildscribe.com/adventure/internal/handler/gin_handler"
	"wildscribe.com/adventure/internal/repository/mongoDB"
	// "wildscribe.com/adventure/internal/routes"
	grpchandler "wildscribe.com/adventure/internal/handler/grpc"
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
	// router := gin.Default()
	log.Println("Connecting to MongoDB")
	db := mongoDB.ConnectDB()
	log.Println("Done!")

	log.Println("Setting up collection")
	repo := mongoDB.NewCollection(db)
	log.Println("Done!")

	log.Println("Setting controller")
	svc := adventure.New(repo)
	log.Println("Done!")
	log.Println("Setting Handler")
	h := grpchandler.New(svc)
	log.Println("Done!")
	log.Println("Setting up route")
	lis, err := net.Listen("tcp", "localhost:8081")
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
	log.Println("Done!")

}

// log.Println("Setting handler")
// handler := grpchandler.New(svc)
// log.Println("Done!")

// log.Println("Setting routes")
// // routes.AdventureRoutes(router, handler)

// log.Println("Done!")

// log.Println("Starting service")
// router.Run(route)
// log.Println("Done! Service is live!")
// }
