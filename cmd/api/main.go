package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/nicholasdly/rest/internal/config"
	"github.com/nicholasdly/rest/internal/db"
	"github.com/nicholasdly/rest/internal/server"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.NewPool(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := server.NewServer(config, db)

	slog.Info(fmt.Sprintf("Listening on %s", config.Address))
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
