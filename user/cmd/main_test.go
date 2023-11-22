package main_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"wildscribe.com/user/internal/controller/user"
	"wildscribe.com/user/internal/handler/gin_handler"
	"wildscribe.com/user/internal/repository/mockdb"
	"wildscribe.com/user/internal/request"

	"wildscribe.com/user/internal/response"
	"wildscribe.com/user/internal/routes"
)

// func TestMain(t *testing.T) {
// 	router := gin.Default()

// 	repo := mockdb.MockCollection()

// 	ctrl := user.New(repo)

// 	handler := gin_handler.New(ctrl)

// 	routes.UserRoute(router, handler)
// }

func TestGetUser(t *testing.T) {
	// Setup
	router := gin.Default()
	repo := mockdb.MockCollection()
	ctrl := user.New(repo)
	handler := gin_handler.New(ctrl)
	routes.UserRoute(router, handler)

	// Set User email and password for POST request
	email, password := "me@gmail.com", "password"

	// Set request and response structs
	var requestBody request.UserRequest
	var responseBody response.UserResponse

	// Assign email/password to requestBody
	requestBody.Data.Type = "user"
	requestBody.Data.Attributes.Email = email
	requestBody.Data.Attributes.Password = password

	//Marshal body as JSON
	body, _ := json.Marshal(requestBody)

	//Setup POST request and response recorder
	req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
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
	name := responseBody.Data.Attributes.Name
	user_id := responseBody.Data.Attributes.User_id

	// Assert that the response code is HTTP 200 (Okay) and fields are correct
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Ian", name)
	assert.Equal(t, "65330eb5fcb829e722f7c40c", user_id)

}
