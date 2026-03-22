package users

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/nicholasdly/rest/internal/common"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll(r.Context())
	if err != nil {
		slog.Error("Failed to retrieve all users.", "error", err)
		http.Error(w, "Failed to retrieve all users.", http.StatusInternalServerError)
		return
	}

	common.RespondJson(w, users, http.StatusOK)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	user, err := h.service.Get(r.Context(), id)
	if errors.Is(err, pgx.ErrNoRows) {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}
	if err != nil {
		slog.Error("Failed to retrieve user.", "id", id, "error", err)
		http.Error(w, "Failed to retrieve user.", http.StatusInternalServerError)
		return
	}

	common.RespondJson(w, user, http.StatusOK)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.Create(r.Context(), req.Username, req.Email)
	if err != nil {
		slog.Error("Failed to create user.", "req", req, "error", err)
		http.Error(w, "Failed to create user.", http.StatusInternalServerError)
		return
	}

	common.RespondJson(w, user, http.StatusCreated)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.Update(r.Context(), id, req.Username, req.Email)
	if errors.Is(err, pgx.ErrNoRows) {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}
	if err != nil {
		slog.Error("Failed to update user.", "req", req, "error", err)
		http.Error(w, "Failed to update user.", http.StatusInternalServerError)
		return
	}

	common.RespondJson(w, user, http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.service.Delete(r.Context(), id)
	if err != nil {
		slog.Error("Failed to delete user.", "id", id, "error", err)
		http.Error(w, "Failed to delete user.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
