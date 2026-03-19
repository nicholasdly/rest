package server

import (
	"net/http"
	"time"

	"github.com/nicholasdly/rest/internal/users"
)

type Server struct {
	server  *http.Server
	handler http.Handler

	userHandler *users.Handler
}

func NewServer() *Server {
	userService := users.NewService()
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
