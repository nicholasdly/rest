package main

import (
	"log"
	"net/http"

	"github.com/nicholasdly/rest/internal/api"
	"github.com/nicholasdly/rest/internal/store"
)

func main() {
	// Initialize dependencies
	userStore := store.NewInMemoryStore()

	// Create API server
	server := api.NewServer(userStore)

	// Start server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", server.Routes()))
}
