package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nicholasdly/rest/internal/db"
	"github.com/nicholasdly/rest/internal/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := db.New("./rest.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := server.New(db)

	errors := make(chan error, 1)
	signals := make(chan os.Signal, 1)

	go func() {
		log.Printf("Server starting on %s", ":8080")
		errors <- server.Start()
	}()

	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errors:
		return fmt.Errorf("server error: %w", err)
	case sig := <-signals:
		log.Printf("Received signal: %v. Starting shutdown...", sig)

		ctx, cancel := context.WithTimeout(
			context.Background(),
			30*time.Second,
		)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}

		log.Println("Server stopped gracefully")
	}

	return nil
}
