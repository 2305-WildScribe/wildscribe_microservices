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
	mongoClient *mongo.Client
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
		log.Println("Error connecting to MongoDB:", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Println("Error connecting to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB")
	return &Database{
		mongoClient: client,
	}
}

// Sets the collection to "adventures"
func NewCollection(database *Database) *Collection {
	return &Collection{
		collection: database.mongoClient.Database("golangAPI").Collection("adventures"),
	}
}

// Set adventure model
var adventure model.Adventure

// Get a single collection from the ID, bind & return adventure model.
func (c *Collection) GetOne(ctx context.Context, id string) (*model.Adventure, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		new_error := fmt.Errorf("MongoDB::GetOne: Decode objId Failed: %w", err)
		return nil, new_error
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	result := c.collection.FindOne(ctx, filter).Decode(&adventure)
	if result != nil {
		new_error := fmt.Errorf("MongoDB::GetOne: FindOne Failed: %w", result)
		return nil, new_error
	}
	return &adventure, nil
}

// Get All Adventures from collection based on a User ID, bind & return []adventure models
func (c *Collection) GetAll(ctx context.Context, id string) ([]*model.Adventure, error) {
	var adventures []*model.Adventure
	filter := bson.D{{Key: "user_id", Value: id}}
	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		new_error := fmt.Errorf("MongoDB::GetAll: Find Failed: %w", err)
		return nil, new_error
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &adventures); err != nil {
		new_error := fmt.Errorf("MongoDB::GetAll: Cursor Failed: %w", err)
		return nil, new_error
	}
	return adventures, nil
}

// Create a new adventure
func (c *Collection) Create(ctx context.Context, adventure *model.Adventure) error {
	result, err := c.collection.InsertOne(ctx, adventure)
	if err != nil {
		new_error := fmt.Errorf("MongoDB::Create: InsertOne Failed: %w", err)
		return new_error
	}
	adventure.Adventure_id = result.InsertedID.(primitive.ObjectID).Hex()
	return err
}

// Update an adventure
func (c *Collection) Update(ctx context.Context, adventure *model.Adventure) (*model.Adventure, error) {
	objId, idErr := primitive.ObjectIDFromHex(adventure.Adventure_id)
	if idErr != nil {
		return nil, fmt.Errorf("MongoDB::Update: Invalid Obj ID: %w", idErr)
	}
	// Set adventure_id to "" to avoid overwrite ID errors.
	adventure.Adventure_id = ""
	filter := bson.D{{Key: "_id", Value: objId}}
	update := bson.D{{Key: "$set", Value: adventure}}

	var updatedAdventure model.Adventure
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := c.collection.FindOneAndUpdate(ctx, filter, update, opts)
	if err := result.Decode(&updatedAdventure); err != nil {
		return nil, fmt.Errorf("MongoDB::Update: FindOneAndUpdate Failed: %w", err)
	}
	return &updatedAdventure, nil
}

// Delete an adventure
func (c *Collection) Delete(ctx context.Context, id string) error {
	objId, idErr := primitive.ObjectIDFromHex(id)
	if idErr != nil {
		return fmt.Errorf("MongoDB::Delete: Invalid Obj ID: %w", idErr)
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	_, err := c.collection.DeleteOne(ctx, filter)
	if err != nil {
		new_error := fmt.Errorf("MongoDB::Delete: DeleteOne failed: %w", err)
		return new_error
	}
	return nil
}
