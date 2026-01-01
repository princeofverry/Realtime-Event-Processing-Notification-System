package handlers

import (
	"encoding/json"
	"event-driven-system/models"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// w untuk ngirim response ke client
// r untuk request dari client (body, header)
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Type      string `json:"type"`
		Message   string `json:"message"`
	}

	// Decoder itu JSON ke GO
	// Encoder GO ke JSON
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	event := models.Event{
		ID:        uuid.NewString(),
		Type:      input.Type,
		Message:   input.Message,
		Status:    "accepted",
		CreatedAt: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	// Encoder GO ke JSON
	json.NewEncoder(w).Encode(event)
}