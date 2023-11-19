package main

import (
	"log"
	"net/http"

	"movieexample.com/rating/internal/controller/rating"
	httphandler "movieexample.com/rating/internal/handler/http"
	"movieexample.com/rating/internal/repository/mongoDB"
)

func main() {
	log.Println("Starting rating service...")
	db := mongoDB.ConnectDB()
	repo := mongoDB.NewCollection(db)
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
