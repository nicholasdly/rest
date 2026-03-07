package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/nicholasdly/rest/internal/models"
)

//
// Mock User Store
//

type MockStore struct {
	users []models.User
}

func (m *MockStore) GetAll() ([]models.User, error) {
	return m.users, nil
}

func (m *MockStore) Get(id int) (models.User, error) {
	if id < len(m.users)-1 || id > len(m.users)-1 {
		return models.User{}, errors.New("user not found")
	}

	return m.users[id], nil
}

func (m *MockStore) Create(user models.User) (models.User, error) {
	user.Id = len(m.users)

	m.users = append(m.users, user)
	return user, nil
}

func (m *MockStore) Update(user models.User) (models.User, error) {
	if user.Id < len(m.users)-1 || user.Id > len(m.users)-1 {
		return models.User{}, errors.New("user not found")
	}

	m.users[user.Id] = user
	return user, nil
}

func (m *MockStore) Delete(id int) error {
	if id < len(m.users)-1 || id > len(m.users)-1 {
		return errors.New("user not found")
	}

	m.users = slices.Delete(m.users, id, id+1)
	return nil
}

//
// Mock Logger
//

type MockLogger struct {
	infos  []string
	errors []string
}

func (m *MockLogger) Info(message string) {
	m.infos = append(m.infos, message)
}

func (m *MockLogger) Error(message string) {
	m.errors = append(m.errors, message)
}

//
// Helpers
//

func setupTestServer() (*Server, *MockStore, *MockLogger) {
	mockStore := &MockStore{
		users: []models.User{
			{Id: 0, FirstName: "John", LastName: "Doe", Email: "hello@email.com"},
			{Id: 1, FirstName: "Jane", LastName: "Doe", Email: "hello@email.com"},
		},
	}
	mockLogger := &MockLogger{}

	server := NewServer(mockStore, mockLogger)
	return server, mockStore, mockLogger
}

//
// Tests
//

func TestHealth(t *testing.T) {
	server, _, _ := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	if w.Body.String() != "OK" {
		t.Errorf("expected body 'OK', got '%s'", w.Body.String())
	}
}

func TestGetUsers(t *testing.T) {
	server, store, _ := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var users []models.User
	if err := json.NewDecoder(w.Body).Decode(&users); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(users) != len(store.users) {
		t.Errorf("expected %d users, got %d", len(store.users), len(users))
	}
}

func TestGetUser(t *testing.T) {
	server, store, _ := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var user models.User
	if err := json.NewDecoder(w.Body).Decode(&user); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if user != store.users[1] {
		t.Errorf("expected %v, got %v", store.users[1], user)
	}
}

func TestCreateUser(t *testing.T) {
	server, store, _ := setupTestServer()

	user := models.User{
		Id:        2,
		FirstName: "Nicholas",
		LastName:  "Ly",
		Email:     "test@email.com",
	}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var created models.User
	if err := json.NewDecoder(w.Body).Decode(&created); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(store.users) != 3 {
		t.Errorf("expected user count of 3, got %d", len(store.users))
	}

	if created != user {
		t.Errorf("expected %v, got %v", user, created)
	}
}

func TestUpdateUser(t *testing.T) {
	server, store, _ := setupTestServer()

	user := models.User{
		Id:        1,
		FirstName: "Nicholas",
		LastName:  "Ly",
		Email:     "test@email.com",
	}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var updated models.User
	if err := json.NewDecoder(w.Body).Decode(&updated); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(store.users) != 2 {
		t.Errorf("expected user count of 2, got %d", len(store.users))
	}

	if updated != user {
		t.Errorf("expected %v, got %v", user, updated)
	}
}

func TestDeleteUser(t *testing.T) {
	server, store, _ := setupTestServer()

	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", w.Code)
	}

	if len(store.users) != 1 {
		t.Errorf("expected user count of 1, got %d", len(store.users))
	}
}
