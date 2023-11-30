package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/go-faker/faker/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"

	adventuremodel "wildscribe.com/adventure/pkg/model"
	usermodel "wildscribe.com/user/pkg/model"
)

func main() {
	var ctx context.Context
	// Setup DB connectio
	db := ConnectDB()
	// Set User Collection
	user_c := NewCollection(db, "golangAPI", "users")
	log.Println("Creating Users")
	adventure_c := NewCollection(db, "golangAPI", "adventures")
	// Create 5000 users and insert them into the array
	var user_models []*usermodel.User
	for i := 0; i < 5000; i++ {
		user_model := CreateUser(i)
		user_models = append(user_models, user_model)
	}
	// Convert []*model.User to []interface{}
	var userInterfaces []interface{}
	for _, u := range user_models {
		userInterfaces = append(userInterfaces, u)
	}
	// Insert all Users into the DB CAN TAKE A FEW MINUTES
	log.Println("Inserting Users")
	user_c.collection.InsertMany(ctx, userInterfaces)
	log.Println("Done!")

	// Get all user IDs from the DB
	log.Println("Fetching Users")
	filter := bson.D{}
	cursor, err := user_c.collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Got Users!")
	var users []string
	// Create User models from return
	for cursor.Next(ctx) {
		var user usermodel.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		// Appends User ID strings to the slice
		users = append(users, user.User_id)
	}
	// test print user IDS
	log.Println(users[:10])
	var adventures []*adventuremodel.Adventure
	// The first 500 users don't have adventures
	// For 500 - 2500 users Create 1 adventures (2000 users)
	for i := 500; i <= 2500 && i < len(users); i++ {
		user := users[i]
		for j := 0; j < 1; j++ {
			adventure := CreateAdventure(user)
			adventures = append(adventures, adventure)
		}
	}
	// For 2501 through 3500 create 10 adventures (1000 users)
	for i := 2501; i <= 3500 && i < len(users); i++ {
		user := users[i]
		for j := 0; j < 10; j++ {
			adventure := CreateAdventure(user)
			adventures = append(adventures, adventure)
		}
	}
	// For 3501 through 4500 create 25 adventures (1000 users)
	for i := 3501; i <= 4500 && i < len(users); i++ {
		user := users[i]
		for j := 0; j < 25; j++ {
			adventure := CreateAdventure(user)
			adventures = append(adventures, adventure)
		}
	}
	// For 4501 through 4800 create 50 adventures (300 users)
	for i := 4501; i <= 4800 && i < len(users); i++ {
		user := users[i]
		for j := 0; j < 50; j++ {
			adventure := CreateAdventure(user)
			adventures = append(adventures, adventure)
		}
	}
	// For 4801 through 4900 create 100 adventures (100 users)
	for i := 4801; i <= 4900 && i < len(users); i++ {
		user := users[i]
		for j := 0; j < 100; j++ {
			adventure := CreateAdventure(user)
			adventures = append(adventures, adventure)
		}
	}
	// for 4901 through 4989 create 1000 adventures (90 users)
	for i := 4901; i <= 4989 && i < len(users); i++ {
		user := users[i]
		for j := 0; j < 1000; j++ {
			adventure := CreateAdventure(user)
			adventures = append(adventures, adventure)
		}
	}
	// For 4990 through 5000 create 10000 adventures (10 users)
	for i := 4990; i <= 4999 && i < len(users); i++ {
		user := users[i]
		for j := 0; j < 10000; j++ {
			adventure := CreateAdventure(user)
			adventures = append(adventures, adventure)
		}
	}
	var adventureInterfaces []interface{}

	// Convert adventure slice to interface slice for mong DB
	for _, a := range adventures {
		adventureInterfaces = append(adventureInterfaces, a)
	}

	log.Println("Adventures Created!")
	log.Println("Inserting Adventures!")
	// Insert into DB. WILL TAKE A FEW MINUTES
	adventure_c.collection.InsertMany(ctx, adventureInterfaces)
	log.Println("Done!")
}

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
	clientOptions := options.Client().ApplyURI("MONGOURIHERE")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB")
	return &Database{
		mongoClient: client,
	}
}

// Sets the collection to "adventures"
func NewCollection(database *Database, db string, collection string) *Collection {
	return &Collection{
		collection: database.mongoClient.Database(db).Collection(collection),
	}
}

func CreateUser(number int) *usermodel.User {
	password := fmt.Sprintf("Password%d", number)
	bytePassword := []byte(password)
	user := &usermodel.User{
		Name:     faker.Name(),
		Email:    fmt.Sprintf("me%d@gmail.com", number),
		Password: HashAndSalt(bytePassword),
	}
	return user
}

func CreateAdventure(user_id string) *adventuremodel.Adventure {
	randomSleep := rand.Intn(24) + 1
	randomDiet := rand.Intn(3501) + 500
	var activities = []string{"Running", "Cycling", "Swimming", "Hiking", "Yoga", "Weightlifting",
		"Aerobics", "Boxing", "Canoeing", "Martial Arts"}
	var stress = []string{"None", "Low", "Moderate", "High", "Max"}
	var hydration = []string{"Dehydrated", "Somewhat Hydrated", "Hydrated", "Very Hydrated"}
	randomActivity := activities[rand.Intn(len(activities))]
	randomStress := stress[rand.Intn(len(stress))]
	randomHydration := hydration[rand.Intn(len(hydration))]
	adventure := &adventuremodel.Adventure{
		User_id:              user_id,
		Activity:             randomActivity,
		Date:                 faker.Date(),
		Image_url:            faker.URL(),
		Stress_level:         randomStress,
		Hours_slept:          int32(randomSleep),
		Sleep_stress_notes:   faker.Sentence(),
		Hydration:            randomHydration,
		Diet:                 int32(randomDiet),
		Diet_hydration_notes: faker.Sentence(),
		Beta_notes:           faker.Sentence(),
		Lat:                  faker.Latitude(),
		Lon:                  faker.Longitude(),
	}

	return adventure
}

func HashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
