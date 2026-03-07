package store

import (
	"testing"

	"github.com/nicholasdly/rest/internal/models"
)

func TestGetAll(t *testing.T) {
	users := map[int]models.User{
		1: {
			Id:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "hello@email.com",
		},
		2: {
			Id:        2,
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "hello@email.com",
		},
	}

	store := NewInMemoryStore()

	store.Lock()
	store.users = users
	store.Unlock()

	result, err := store.GetAll()
	if err != nil || len(result) != 2 {
		t.Errorf("failed to retrieve all users")
	}
}

func TestGet(t *testing.T) {
	user := models.User{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "hello@email.com",
	}

	store := NewInMemoryStore()

	store.Lock()
	store.users[1] = user
	store.Unlock()

	result, err := store.Get(1)
	if err != nil || result != user {
		t.Errorf("failed to retrieve user")
	}
}

func TestCreate(t *testing.T) {
	store := NewInMemoryStore()

	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "hello@email.com",
	}

	_, err := store.Create(user)
	if err != nil || len(store.users) != 1 {
		t.Errorf("failed to create user")
	}
}

func TestUpdate(t *testing.T) {
	user := models.User{
		Id:        1,
		FirstName: "Nicholas",
		LastName:  "Ly",
		Email:     "greetings@email.com",
	}

	store := NewInMemoryStore()

	store.Lock()
	store.users[1] = models.User{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "hello@email.com",
	}
	store.Unlock()

	result, err := store.Update(user)
	if err != nil || result != user {
		t.Errorf("failed to update user")
	}
}

func TestDelete(t *testing.T) {
	users := map[int]models.User{
		1: {
			Id:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "hello@email.com",
		},
		2: {
			Id:        2,
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "hello@email.com",
		},
	}

	store := NewInMemoryStore()

	store.Lock()
	store.users = users
	store.Unlock()

	err := store.Delete(1)
	if err != nil || len(store.users) != 1 {
		t.Errorf("failed to delete users")
	}
}
