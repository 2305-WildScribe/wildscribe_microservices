package main_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"

// 	"wildscribe.com/adventure/internal/controller/adventure"
// 	"wildscribe.com/adventure/internal/handler/gin_handler"
// 	"wildscribe.com/adventure/internal/repository/mockDB"
// 	"wildscribe.com/adventure/internal/response"
// 	"wildscribe.com/adventure/internal/routes"
// )

// var router *gin.Engine

// func TestMain(m *testing.M) {
// 	// Setup
// 	router = gin.Default()
// 	repo := mockDB.MockCollection()
// 	ctrl := adventure.New(repo)
// 	handler := gin_handler.New(ctrl)
// 	routes.AdventureRoutes(router, handler)

// 	// Run tests
// 	exitCode := m.Run()
// 	// Exit with the exit code from the tests
// 	os.Exit(exitCode)
// }

// func makeRequest(t *testing.T, method, path string, body interface{}) *httptest.ResponseRecorder {
// 	bodyJSON, err := json.Marshal(body)
// 	if err != nil {
// 		t.Fatalf("Failed to marshal request body: %v", err)
// 	}
// 	req, err := http.NewRequest(method, path, bytes.NewBuffer(bodyJSON))
// 	if err != nil {
// 		t.Fatalf("Failed to create request: %v", err)
// 	}
// 	response := httptest.NewRecorder()
// 	router.ServeHTTP(response, req)
// 	return response
// }

// func unmarshalResponseBody(t *testing.T, response *httptest.ResponseRecorder, responseBody interface{}) {
// 	err := json.Unmarshal(response.Body.Bytes(), responseBody)
// 	if err != nil {
// 		log.Println(response.Body)
// 		t.Fatalf("Failed to unmarshal response body: %v", err)
// 	}
// }

// func TestGetAdventure(t *testing.T) {
// 	userID := "652edaa67a75034ea37c6652"
// 	adventureID := "656001daf827a04b7a66bafa"
// 	requestBody := map[string]interface{}{
// 		"data": map[string]interface{}{
// 			"type": "adventure",
// 			"attributes": map[string]interface{}{
// 				"adventure_id": adventureID,
// 				"user_id":      userID,
// 			},
// 		},
// 	}
// 	var responseBody response.AdventureResponse

// 	response := makeRequest(t, http.MethodPost, "/adventure", requestBody)
// 	unmarshalResponseBody(t, response, &responseBody)

// 	assert.Equal(t, http.StatusOK, response.Code)
// 	assert.Equal(t, adventureID, responseBody.Data.Attributes[0].Adventure_id)
// 	assert.Equal(t, userID, responseBody.Data.Attributes[0].User_id)
// 	assert.Equal(t, "Test", responseBody.Data.Attributes[0].Activity)
// }

// func TestCreateAdventure(t *testing.T) {
// 	userID := "652edaa67a75034ea37c6652"
// 	requestBody := map[string]interface{}{
// 		"data": map[string]interface{}{
// 			"type": "adventure",
// 			"attributes": map[string]interface{}{
// 				"user_id":  userID,
// 				"activity": "Test",
// 			},
// 		},
// 	}
// 	var responseBody response.AdventureResponse

// 	response := makeRequest(t, http.MethodPost, "/create_adventure", requestBody)
// 	unmarshalResponseBody(t, response, &responseBody)

// 	assert.Equal(t, http.StatusOK, response.Code)
// 	assert.Equal(t, "656001daf827a04b7a66bafa", responseBody.Data.Attributes[0].Adventure_id)
// 	assert.Equal(t, userID, responseBody.Data.Attributes[0].User_id)
// 	assert.Equal(t, "Test", responseBody.Data.Attributes[0].Activity)
// }
