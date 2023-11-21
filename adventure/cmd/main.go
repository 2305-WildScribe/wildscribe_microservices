package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"wildscribe.com/adventure/internal/controller/adventure"
	"wildscribe.com/adventure/internal/handler/gin_handler"
	"wildscribe.com/adventure/internal/repository/mongoDB"
	"wildscribe.com/adventure/internal/routes"
)

func main() {
	var port string
	log.Println("Starting wildscribe adventure service...")
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
		port = "8082"
	}

	route := fmt.Sprintf("0.0.0.0:%s", port)

	router := gin.Default()

	db := mongoDB.ConnectDB()

	repo := mongoDB.NewCollection(db)

	ctrl := adventure.New(repo)

	handler := gin_handler.New(ctrl)

	routes.AdventureRoutes(router, handler)

	router.Run(route)
}
