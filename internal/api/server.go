package api

import (
	"net/http"

	"github.com/nicholasdly/rest/internal/store"
)

type Server struct {
	store store.UserStore
}

func NewServer(store store.UserStore) *Server {
	return &Server{
		store: store,
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

	return mux
}
