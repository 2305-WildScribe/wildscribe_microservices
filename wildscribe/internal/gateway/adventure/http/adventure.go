package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"wildscribe.com/wildscribe/internal/request"
	"wildscribe.com/wildscribe/pkg/model"
)

// Gateway defines movie metadata HTTP gateway.
type Gateway struct {
	addr string
}

// New creates a new movie metadata HTTP gateway.
func New(addr string) *Gateway {
	return &Gateway{addr}
}

// Get gets movie metadata for a given movie ID.

func (g *Gateway) GetAdventure(ctx context.Context, gateway_request request.AdventureRequest) (*model.Adventure, error) {
	var adventureResponse request.AdventureRequest
	// Convert the data map to JSON
	jsonData, err := json.Marshal(gateway_request)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP request with the request body
	req, err := http.NewRequest(http.MethodPost, g.addr+"/adventure", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		log.Println("Error:", err)
		return nil, err
	}
	log.Printf("Raw JSON Response: %s", body)

	if err := json.Unmarshal(body, &adventureResponse); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return nil, err
	}

	adventure := model.NewAdventure(adventureResponse)
	return adventure, nil
}
