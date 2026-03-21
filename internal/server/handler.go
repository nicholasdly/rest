package server

import (
	"net/http"

	"github.com/nicholasdly/rest/internal/middleware"
)

func (s *Server) setupHandler() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handleHealth)

	mux.HandleFunc("GET /users", s.userHandler.GetAll)
	mux.HandleFunc("GET /users/{id}", s.userHandler.Get)
	mux.HandleFunc("POST /users", s.userHandler.Create)
	mux.HandleFunc("PUT /users/{id}", s.userHandler.Update)
	mux.HandleFunc("DELETE /users/{id}", s.userHandler.Delete)

	s.handler = mux
	s.handler = middleware.AuthMiddleware(s.config.ApiKey, s.handler)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}
