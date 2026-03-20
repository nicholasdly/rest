package main

import (
	"log"

	"github.com/nicholasdly/rest/internal/config"
	"github.com/nicholasdly/rest/internal/server"
)

func main() {
	config := config.Load()

	server := server.NewServer(config)

	log.Print("Listening on :8080")
	log.Fatal(server.Start())
}
