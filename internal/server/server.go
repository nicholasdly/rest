package server

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/nicholasdly/rest/internal/users"
)

type Server struct {
	server  *http.Server
	handler http.Handler

	userHandler *users.Handler
}

func New(db *sql.DB) *Server {
	// Initialize repositories
	userRepo := users.NewRepository(db)

	// Initialize services
	userService := users.NewService(userRepo)

	// Initialize handlers
	userHandler := users.NewHandler(userService)

	s := &Server{
		userHandler: userHandler,
	}

	s.setupHandler()

	s.server = &http.Server{
		Addr:         ":8080",
		Handler:      s.handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return s
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
