package mongoDB

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"movieexample.com/rating/internal/config"
	"movieexample.com/rating/pkg/model"
)

// Set Database stuct
type Database struct {
	db *mongo.Client
}

// Set Collection struct
type Collection struct {
	collection *mongo.Collection
}

// Connects to MongoDB
func ConnectDB() *Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.EnvMongoURI()))
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

// Sets the collection to "metadata"
func NewCollection(db *Database) *Collection {
	return &Collection{
		collection: db.db.Database("micro_service_db").Collection("rating"),
	}
}

// Set metadata model
var metadata model.Rating

// Get a single collection from the ID, bind & return metadata model.
func (c *Collection) Get(ctx context.Context, id string) (*model.Rating, error) {
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
