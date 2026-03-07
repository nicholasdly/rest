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

func (server *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", server.health)
	mux.HandleFunc("GET /users", server.getUsers)
	mux.HandleFunc("GET /users/{id}", server.getUser)
	mux.HandleFunc("POST /users", server.createUser)

	return mux
}
