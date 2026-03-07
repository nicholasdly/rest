package main

import (
	"log"
	"net/http"

	"github.com/nicholasdly/rest/internal/api"
	"github.com/nicholasdly/rest/internal/logger"
	"github.com/nicholasdly/rest/internal/store"
)

func main() {
	// Initialize dependencies
	userStore := store.NewInMemoryStore()
	logger := logger.NewStdLogger()

	// Create API server
	server := api.NewServer(userStore, logger)

	// Start server
	logger.Info("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", server.Handler()))
}
