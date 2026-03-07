package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nicholasdly/rest/internal/models"
)

func (s *Server) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) getUsers(w http.ResponseWriter, req *http.Request) {
	users, err := s.store.GetAll()

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondJson(w, http.StatusOK, users)
}

func (s *Server) getUser(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	user, err := s.store.Get(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondJson(w, http.StatusOK, user)
}

func (s *Server) createUser(w http.ResponseWriter, req *http.Request) {
	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := s.store.Create(user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondJson(w, http.StatusOK, created)
}

func (s *Server) updateUser(w http.ResponseWriter, req *http.Request) {
	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := s.store.Get(user.Id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	updated, err := s.store.Update(user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondJson(w, http.StatusOK, updated)
}

func (s *Server) deleteUser(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	if _, err := s.store.Get(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := s.store.Delete(id); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
