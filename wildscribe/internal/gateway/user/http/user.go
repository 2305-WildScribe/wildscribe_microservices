package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"wildscribe.com/wildscribe/internal/gateway"
	"wildscribe.com/wildscribe/internal/request"
	"wildscribe.com/wildscribe/pkg/model"
)

// Gateway defines movie rating HTTP gateway.
type Gateway struct {
	addr string
}

// New creates a new movie rating HTTP gateway.
func New(addr string) *Gateway {
	return &Gateway{addr}
}

func (g *Gateway) GetAggregatedRating(ctx context.Context, request request.UserRequest) (*model.User, error) {
	var user model.User
	req, err := http.NewRequest(http.MethodPost, g.addr+"/user", nil)
	if err != nil {
		return &user, err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", string(recordID))
	values.Add("type", fmt.Sprintf("%v", recordType))
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return 0, gateway.ErrNotFound
	} else if resp.StatusCode/100 != 2 {
		return 0, fmt.Errorf("non-2xx response: %v", resp)
	}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return 0, err
	}
	return &user, nil
}
