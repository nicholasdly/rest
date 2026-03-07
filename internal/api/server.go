package api

import (
	"net/http"

	"github.com/nicholasdly/rest/internal/logger"
	"github.com/nicholasdly/rest/internal/store"
)

type Server struct {
	store  store.UserStore
	logger logger.Logger
}

func NewServer(store store.UserStore, logger logger.Logger) *Server {
	return &Server{
		store:  store,
		logger: logger,
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", s.health)
	mux.HandleFunc("GET /users", s.getUsers)
	mux.HandleFunc("GET /users/{id}", s.getUser)
	mux.HandleFunc("POST /users", s.createUser)
	mux.HandleFunc("PUT /users", s.updateUser)
	mux.HandleFunc("DELETE /users/{id}", s.deleteUser)

	return s.loggingMiddleware(mux)
}
