package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"wildscribe.com/adventure/internal/controller/adventurecontroller"
	"wildscribe.com/adventure/internal/repository"
)

// Handler defines a movie adventure HTTP handler.
type Handler struct {
	ctrl *adventurecontroller.Controller
}

// New creates a new movie adventure HTTP handler.
func New(ctrl *adventurecontroller.Controller) *Handler {
	return &Handler{ctrl}
}

// GetAdventure handles GET /adventure requests.
func (h *Handler) GetAdventure(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := req.Context()
	adv, err := h.ctrl.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(adv); err != nil {
		log.Printf("JSON encode error: %v\n", err)
	}
}
