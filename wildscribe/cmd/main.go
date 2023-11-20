package main

import (
	"log"
	"net/http"

	"wildscribe.com/wildscribe/internal/controller/user"
	adventuregateway "wildscribe.com/wildscribe/internal/gateway/adventure/http"
	ratinggateway "wildscribe.com/wildscribe/internal/gateway/user/http"
	httphandler "wildscribe.com/wildscribe/internal/handler/http"
)

func main() {
	log.Println("Starting movie service...")
	adventureGateway := adventuregateway.New("http://localhost:8081")
	log.Println("Connecting to Metadata service")
	ratingGateway := ratinggateway.New("http://localhost:8082")
	log.Println("Connecting to Rating service")
	ctrl := movie.New(ratingGateway, adventureGateway)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
