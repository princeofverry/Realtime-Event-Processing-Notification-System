package routes

import (
	"event-driven-system/api/handlers"
	"net/http"
)

func Register() http.Handler {
	// router dari go
	mux := http.NewServeMux()
	mux.HandleFunc("/events", handlers.CreateEvent)
	return mux
}