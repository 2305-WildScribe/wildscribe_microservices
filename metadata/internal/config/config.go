package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Access Env for MONGOURI
func EnvMongoURI() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("error: %v\n", err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGOURI")
}
