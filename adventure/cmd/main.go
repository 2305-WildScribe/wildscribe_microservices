package main

import (
	"log"
	"net/http"

	"movieexample.com/adventure/internal/controller/adventure"
	httphandler "movieexample.com/adventure/internal/handler/http"
	"movieexample.com/adventure/internal/repository/mongoDB"
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

	ctrl := adventure.New(repo)

	h := httphandler.New(ctrl)

	http.Handle("/adventure", http.HandlerFunc(h.GetAdventure))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("Error starting HTTP server:", err)
	}
}