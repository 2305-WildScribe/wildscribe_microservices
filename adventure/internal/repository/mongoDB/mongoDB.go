package mongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"wildscribe.com/adventure/internal/config"
	"wildscribe.com/adventure/pkg/model"
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
		collection: db.db.Database("golangAPI").Collection("adventures"),
	}
}

// Set adventure model
var adventure model.Adventure

// Get a single collection from the ID, bind & return adventure model.
func (c *Collection) GetOne(ctx context.Context, id string) (*model.Adventure, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objId}}
	result := c.collection.FindOne(ctx, filter).Decode(&adventure)
	if result != nil {
		log.Printf("MongoDB error: %v\n", err)
		return nil, err
	}
	return &adventure, nil
}

// Create a new adventure
func (c *Collection) Create(ctx context.Context, adventure *model.Adventure) error {
	_, err := c.collection.InsertOne(ctx, adventure)
	if err != nil {
			log.Printf("MongoDB error: %v\n", err)
			return err
	}
	return nil
}

// Update an adventure
func (c *Collection) Update(ctx context.Context, updatedAdventure *model.Adventure) error {
	filter := bson.D{{Key: "_id", Value: updatedAdventure.Adventure_id}}
	update := bson.D{{Key: "$set", Value: updatedAdventure}}
	_, err := c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
			log.Printf("MongoDB error: %v\n", err)
			return err
	}
	return nil
}

// Delete an adventure
func (c *Collection) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Failed to convert ID to ObjectID: %v\n", err)
		return err
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err = c.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("MongoDB error: %v\n", err)
		return err
	}
	return nil
}

