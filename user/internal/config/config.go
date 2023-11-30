package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Access Env for MONGOURI
func EnvMongoURI() string {
	env := os.Getenv("ENV")
	if env == "PROD" {
		uri := os.Getenv("MONGOURI")
		log.Println(uri)
		return uri
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return os.Getenv("MONGOURI")
}

func EnvMongoDB() string {
	env := os.Getenv("ENV")
	if env == "PROD" {
		return os.Getenv("DATABASE")
	}
	return "golangAPI"
}

func EnvMongoColleciton() string {
	env := os.Getenv("ENV")
	if env == "PROD" {
		return os.Getenv("COLLECTION")
	}
	return "users"
}
