package main_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"wildscribe.com/adventure/internal/controller/adventure"
	"wildscribe.com/adventure/internal/handler/gin_handler"
	"wildscribe.com/adventure/internal/repository/mockDB"
	"wildscribe.com/adventure/internal/request"

	"wildscribe.com/adventure/internal/response"
	"wildscribe.com/adventure/internal/routes"
)

func TestGetAdventure(t *testing.T) {
	// Setup
	router := gin.Default()
	repo := mockDB.MockCollection()
	ctrl := adventure.New(repo)
	handler := gin_handler.New(ctrl)
	routes.AdventureRoutes(router, handler)

	// Set User email and password for POST request
	id := "656001daf827a04b7a66bafa"

	// Set request and response structs
	var requestBody request.AdventureRequest
	var responseBody response.AdventureResponse

	// Assign email/password to requestBody
	requestBody.Data.Type = "adventure"
	requestBody.Data.Attributes.Adventure_id = id

	//Marshal body as JSON
	body, _ := json.Marshal(requestBody)

	//Setup POST request and response recorder
	req, _ := http.NewRequest(http.MethodPost, "/adventure", bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	// Send reqeusts
	router.ServeHTTP(response, req)

	// Unmarshal JSON body of response and bind it to responseBody
	err := json.Unmarshal(response.Body.Bytes(), &responseBody)
	if err != nil {
		log.Println(response.Body)
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	// Access the Name and User_id from responseBody
	adventure_id := responseBody.Data.Attributes[0].Adventure_id
	user_id := responseBody.Data.Attributes[0].User_id
	activity := responseBody.Data.Attributes[0].Activity

	// Assert that the response code is HTTP 200 (Okay) and fields are correct
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, id, adventure_id)
	assert.Equal(t, "652edaa67a75034ea37c6652", user_id)
	assert.Equal(t, "Test", activity)
}

func TestGetAllAdventures(t *testing.T) {
	// Setup
	router := gin.Default()
	repo := mockDB.MockCollection()
	ctrl := adventure.New(repo)
	handler := gin_handler.New(ctrl)
	routes.AdventureRoutes(router, handler)

	// Set User email and password for POST request
	id := "652edaa67a75034ea37c6652"

	// Set request and response structs
	var requestBody request.AdventureRequest
	var responseBody response.AdventureResponse

	// Assign email/password to requestBody
	requestBody.Data.Type = "adventure"
	requestBody.Data.Attributes.Adventure_id = id

	//Marshal body as JSON
	body, _ := json.Marshal(requestBody)

	//Setup POST request and response recorder
	req, _ := http.NewRequest(http.MethodPost, "/adventures", bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	// Send reqeusts
	router.ServeHTTP(response, req)

	// Unmarshal JSON body of response and bind it to responseBody
	err := json.Unmarshal(response.Body.Bytes(), &responseBody)
	if err != nil {
		log.Println(response.Body)
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	// Access the Name and User_id from responseBody
	adventure_id := responseBody.Data.Attributes[0].Adventure_id
	user_id := responseBody.Data.Attributes[0].User_id
	activity := responseBody.Data.Attributes[0].Activity

	// Assert that the response code is HTTP 200 (Okay) and fields are correct
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, id, adventure_id)
	assert.Equal(t, "652edaa67a75034ea37c6652", user_id)
	assert.Equal(t, "Test", activity)
}
