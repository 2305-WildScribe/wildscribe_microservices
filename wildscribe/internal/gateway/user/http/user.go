package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"wildscribe.com/user/pkg/model"
	"wildscribe.com/wildscribe/internal/gateway"
	"wildscribe.com/wildscribe/internal/request"
)

// Gateway defines the User microserive HTTP gateway.
type Gateway struct {
	addr string
}

// New creates a new User microservice HTTP gateway.
func New(addr string) *Gateway {
	return &Gateway{addr}
}

func (g *Gateway) GetUser(ctx context.Context, request request.UserRequest) (*model.User, error) {
	var user model.User

	// Convert the User request to JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP request with the request body
	req, err := http.NewRequest(http.MethodPost, g.addr+"/user", bytes.NewBuffer(jsonData))
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

	if resp.StatusCode == http.StatusNotFound {
		return nil, gateway.ErrNotFound
	} else if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("non-2xx response: %v", resp)
	}

	// Decode the response body
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
