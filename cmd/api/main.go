package main

import (
	"log"

	"github.com/nicholasdly/rest/internal/server"
)

func main() {
	server := server.NewServer()

	log.Print("Listening on :8080")
	log.Fatal(server.Start())
}
