package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nicholasdly/rest/internal/models"
)

func (server *Server) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (server *Server) getUsers(w http.ResponseWriter, req *http.Request) {
	users, err := server.store.GetAll()

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondJson(w, http.StatusOK, users)
}

func (server *Server) getUser(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	user, err := server.store.Get(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondJson(w, http.StatusOK, user)
}

func (server *Server) createUser(w http.ResponseWriter, req *http.Request) {
	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := server.store.Create(user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondJson(w, http.StatusOK, created)
}

func (server *Server) updateUser(w http.ResponseWriter, req *http.Request) {
	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := server.store.Get(user.Id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	updated, err := server.store.Update(user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondJson(w, http.StatusOK, updated)
}

func (server *Server) deleteUser(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	if _, err := server.store.Get(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := server.store.Delete(id); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
