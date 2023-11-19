package mongoDB

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"wildscribe.com/user/internal/config"
	"wildscribe.com/user/pkg/model"
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
	clientOptions := options.Client().ApplyURI(config.EnvMongoURI())

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
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
		collection: db.db.Database("golangAPI").Collection("users"),
	}
}

// Set user model
var user model.User

// Get a single collection from the ID, bind & return user model.
func (c *Collection) Get(ctx context.Context, email string) (*model.User, error) {

	filter := bson.D{{Key: "email", Value: email}}
	err := c.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Printf("MongoDB error: %v\n", err)
		return nil, err
	}
	return &user, nil
}
