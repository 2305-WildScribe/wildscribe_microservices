package main

import (
	"log"
	"net/http"

	"wildscribe.com/adventure/internal/controller/adventurecontroller"
	httphandler "wildscribe.com/adventure/internal/handler/http"
	"wildscribe.com/adventure/internal/repository/mongoDB"
)

func main() {
	log.Println("Starting adventure service...")

	db, err := mongoDB.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	log.Println("Connected to MongoDB")

	repo := mongoDB.NewCollection(db)
	if repo == nil {
		log.Fatal("Failed to connect to Collection")
	}

	log.Println("Connected to Collection")

	ctrl := adventurecontroller.New(repo)

	h := httphandler.New(ctrl)

	http.Handle("/adventure", http.HandlerFunc(h.GetAdventure))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("Error starting HTTP server:", err)
	}
}
