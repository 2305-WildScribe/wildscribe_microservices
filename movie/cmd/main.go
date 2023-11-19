package main

import (
	"log"
	"net/http"

	"movieexample.com/movie/internal/controller/movie"
	metadatagateway "movieexample.com/movie/internal/gateway/metadata/http"
	ratinggateway "movieexample.com/movie/internal/gateway/rating/http"
	httphandler "movieexample.com/movie/internal/handler/http"
)

func main() {
	log.Println("Starting movie service...")
	metadataGateway := metadatagateway.New("http://localhost:8081")
	log.Println("Connecting to Metadata service")
	ratingGateway := ratinggateway.New("http://localhost:8082")
	log.Println("Connecting to Rating service")
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
