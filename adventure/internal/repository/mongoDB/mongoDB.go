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
	"movieexample.com/adventure/internal/config"
	"movieexample.com/adventure/pkg/model"
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
func ConnectDB() (*Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB")
	return &Database{
		db: client,
	}, nil
}

// Sets the collection to "adventure"
func NewCollection(db *Database) *Collection {
	return &Collection{
		collection: db.db.Database("micro_service_db").Collection("adventure"),
	}
}

// Set adventure model
var adventure model.Adventure

// Get a single collection from the ID, bind & return adventure model.
func (c *Collection) Get(ctx context.Context, id string) (*model.Adventure, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ObjID: %v\n", err)
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	result := c.collection.FindOne(ctx, filter).Decode(&adventure)
	if result != nil {
		log.Printf("MongoDB error: %v\n", err)
	}
	return &adventure, err
}
