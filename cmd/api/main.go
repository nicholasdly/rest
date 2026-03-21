package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/nicholasdly/rest/internal/config"
	"github.com/nicholasdly/rest/internal/server"
)

func main() {
	config, err := config.Load()
	if err != nil {
		slog.Error("Failed to load environment configuration.", "err", err)
		os.Exit(1)
	}

	server := server.NewServer(&config)

	slog.Info(fmt.Sprintf("Listening on %s", config.Address))
	if err := server.Start(); err != nil {
		slog.Error("Server stopped.", "err", err)
		os.Exit(1)
	}
}
