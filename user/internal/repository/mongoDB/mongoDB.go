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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// Create a new user
func (c *Collection) Create(ctx context.Context, user *model.User) error {
	result, err := c.collection.InsertOne(ctx, user)
	if err != nil {
		new_error := fmt.Errorf("MongoDB::Create: InsertOne Failed: %w", err)
		return new_error
	}
	user.User_id = result.InsertedID.(primitive.ObjectID).Hex()
	return err
}

// Update an user
func (c *Collection) Update(ctx context.Context, user *model.User) error {

	filter := bson.D{{Key: "_id", Value: user.User_id}}
	update := bson.D{{Key: "$set", Value: user}}
	_, err := c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		new_error := fmt.Errorf("MongoDB::Update: UpdateOne Failed: %w", err)
		return new_error
	}

	return nil
}

// Delete an user
func (c *Collection) Delete(ctx context.Context, id string) error {
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := c.collection.DeleteOne(ctx, filter)
	if err != nil {
		new_error := fmt.Errorf("MongoDB::Delete: DeleteOne failed: %w", err)
		return new_error
	}
	return nil
}
