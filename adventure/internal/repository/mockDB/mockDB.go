package mockDB

import (
	"context"
	"fmt"
	"log"

	"github.com/mjarkk/mongomock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"wildscribe.com/adventure/pkg/model"
)

// This is a mock of MongoDB using mongomock.
// Set Collection struct
type Collection struct {
	collection *mongomock.Collection
}

// Setups a MockDB using mongomock
func SetupMockDB(collection string) *Collection {
	client := mongomock.NewDB().Collection(collection)

	fmt.Println("Connected to MongoDB")
	return &Collection{
		collection: client,
	}
}

// Set adventure model

// Get a single collection from the ID, bind & return adventure model.
func (c *Collection) GetOne(ctx context.Context, id string) (*model.Adventure, error) {
	adventure := model.Adventure{}
	// Checks objId
	objId, idErr := primitive.ObjectIDFromHex(id)
	if idErr != nil {
		return nil, fmt.Errorf("MongoDB::GetOne: Decode objId Failed: %w", idErr)
	}
	filter := bson.M{"_id": objId}

	err := c.collection.FindFirst(&adventure, filter)
	if err != nil {
		return nil, fmt.Errorf("MongoDB::GetOne: FindOne Failed: %w", err)
	}

	return &adventure, nil
}

// Get All Adventures from collection based on a User ID, bind & return []adventure models
func (c *Collection) GetAll(ctx context.Context, id string) ([]*model.Adventure, error) {
	var adventures []*model.Adventure
	filter := bson.M{"user_id": id}
	cursor, err := c.collection.FindCursor(filter)
	if err != nil {
		return nil, fmt.Errorf("MongoDB::GetAll: Find Failed: %w", err)
	}
	for cursor.Next() {
		adventure := &model.Adventure{}
		err := cursor.Decode(&adventure)
		if err != nil {
			log.Fatal(err)
		}
		adventures = append(adventures, adventure)
	}
	return adventures, nil
}

// Create a new adventure
func (c *Collection) Create(ctx context.Context, adventure *model.Adventure) error {
	objectID := primitive.NewObjectID()
	document := bson.M{
		"_id":      objectID,
		"user_id":  adventure.User_id,
		"Activity": adventure.Activity,
	}
	err := c.collection.Insert(document)
	if err != nil {
		return fmt.Errorf("MongoDB::Create: InsertOne Failed: %w", err)
	}
	adventure.Adventure_id = objectID.Hex()

	return err
}

// Update an adventure
func (c *Collection) Update(ctx context.Context, adventure *model.Adventure) (*model.Adventure, error) {
	objId, idErr := primitive.ObjectIDFromHex(adventure.Adventure_id)
	if idErr != nil {
		return nil, fmt.Errorf("MongoDB::Update: Invalid Obj ID: %w", idErr)
	}
	err := c.collection.ReplaceFirstByID(objId, adventure)
	if err != nil {
		return nil, fmt.Errorf("MongoDB::Update: FindOneAndUpdate Failed: %w", err)
	}
	return adventure, nil
}

// Delete an adventure
func (c *Collection) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	err := c.collection.Delete(filter)
	if err != nil {
		return fmt.Errorf("MongoDB::Delete: DeleteOne failed: %w", err)
	}
	return nil
}
