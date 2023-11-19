package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"wildscribe.com/user/internal/controller/user"
	"wildscribe.com/user/internal/handler/gin_handler"
	"wildscribe.com/user/internal/repository/mongoDB"
	"wildscribe.com/user/internal/routes"
)

func main() {

	router := gin.Default()
	log.Println("Starting rating service...")

	db := mongoDB.ConnectDB()

	repo := mongoDB.NewCollection(db)

	ctrl := user.New(repo)

	handler := gin_handler.New(ctrl)

	routes.UserRoute(router, handler)

	router.Run("localhost:6000")
}
