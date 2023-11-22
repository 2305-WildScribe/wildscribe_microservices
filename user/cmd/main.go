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
	log.Println("Starting wildscribe user service...")
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
		port = "8081"
	}

	route := fmt.Sprintf("0.0.0.0:%s", port)

	router := gin.Default()

	db := mongoDB.ConnectDB()

	repo := mongoDB.NewCollection(db)

	ctrl := user.New(repo)

	handler := gin_handler.New(ctrl)

	routes.UserRoute(router, handler)

	router.Run(route)
}
