package server

import (
	"net/http"

	"github.com/nicholasdly/rest/internal/config"
	"github.com/nicholasdly/rest/internal/users"
)

type Server struct {
	server  *http.Server
	handler http.Handler
	config  *config.Config

	userHandler *users.Handler
}

func NewServer(config *config.Config) *Server {
	userService := users.NewService()
	userHandler := users.NewHandler(userService)

	s := &Server{
		config:      config,
		userHandler: userHandler,
	}

	s.setupHandler()

	s.server = &http.Server{
		Addr:         config.Address,
		Handler:      s.handler,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}

	return s
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
