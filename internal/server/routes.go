package server

import "net/http"

func (s *Server) setupHandler() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", s.handleHealth)

	mux.HandleFunc("GET /users", s.userHandler.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", s.userHandler.GetUser)
	mux.HandleFunc("POST /users", s.userHandler.CreateUser)
	mux.HandleFunc("PUT /users/{id}", s.userHandler.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", s.userHandler.DeleteUser)

	s.handler = mux
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
