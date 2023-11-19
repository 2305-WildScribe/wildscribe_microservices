package main

import (
	"log"
	"net/http"

	"movieexample.com/metadata/internal/controller/metadata"
	httphandler "movieexample.com/metadata/internal/handler/http"
	"movieexample.com/metadata/internal/repository/mongoDB"
)

func main() {
	log.Println("Starting metadata service...")
	db := mongoDB.ConnectDB()
	log.Println("Connected to MongoDB")
	repo := mongoDB.NewCollection(db)
	log.Println("Connected to Collection")
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
