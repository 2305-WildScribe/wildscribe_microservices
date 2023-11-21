package http

import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// "wildscribe.com/user/pkg/model"
// "wildscribe.com/wildscribe/internal/gateway"
)

// Gateway defines movie rating HTTP gateway.
type Gateway struct {
	addr string
}

// New creates a new movie rating HTTP gateway.
func New(addr string) *Gateway {
	return &Gateway{addr}
}
