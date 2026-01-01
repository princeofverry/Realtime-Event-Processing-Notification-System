package main

import (
	"event-driven-system/api/routes"
	"log"
	"net/http"
)

func main() {
	handler := routes.Register()
	
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}