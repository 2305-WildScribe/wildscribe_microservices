package mongoDB

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"movieexample.com/metadata/pkg/model"
)

func EnvMongoURI() string {
	if os.Getenv("PROD_ENV") == "production" {
		return os.Getenv("MONGOURI")
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv("MONGOURI")
	}
}

type Database struct {
	db *mongo.Client
}

func NewDatabase() *Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return &Database{
		db: client,
	}
}

type Collection struct {
	collection *mongo.Collection
}

func NewCollection(db *Database) *Collection {
	return &Collection{
		collection: db.db.Database("micro_service_db").Collection("metadata"),
	}
}

var metadata model.Metadata

func (c *Collection) Get(ctx context.Context, id string) (*model.Metadata, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ObjID: %v\n", err)
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	result := c.collection.FindOne(ctx, filter).Decode(&metadata)
	if result != nil {
		log.Printf("MongoDB error: %v\n", err)
	}
	return &metadata, err
}
