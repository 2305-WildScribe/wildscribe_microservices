package main

import (
	"log"
	// "net/http"
	"fmt"
	"os"
	"wildscribe.com/wildscribe/internal/controller/user"
	"wildscribe.com/wildscribe/internal/routes"

	"github.com/gin-gonic/gin"

	// adventuregateway "wildscribe.com/wildscribe/internal/gateway/adventure/http"
	usergateway "wildscribe.com/wildscribe/internal/gateway/user/http"
	ginhandler "wildscribe.com/wildscribe/internal/handler/gin_handler"
)

func main() {
	var port string
	log.Println("Starting Wildscribe Application...")
	env := os.Getenv("ENV")

	if env == "PROD" {
		port = os.Getenv("PORT")
	} else {

		// f, err := os.Open("configs/base.yml")

		// if err != nil {
		// 	panic(err)
		// }

		// defer f.Close()

		// var cfg ServiceConfig

		// if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		// 	panic(err)
		// }
		port = "8080"
	}

	route := fmt.Sprintf("0.0.0.0:%s", port)

	router := gin.Default()

	log.Println("Connecting to Wildscribe adventure microservice service")
	// adventureGateway := adventuregateway.New("http://localhost:8081")
	log.Println("Connecting to Wildscribe user microservice")
	userGateway := usergateway.New("http://0.0.0.0:8081")
	ctrl := user.New(userGateway)

	handler := ginhandler.New(ctrl)

	routes.UserRoute(router, handler)

	router.Run(route)
}
