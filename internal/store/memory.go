package store

import (
	"fmt"
	"sync"

	"github.com/nicholasdly/rest/internal/models"
)

type InMemoryStore struct {
	sync.Mutex

	users  map[int]models.User
	nextId int
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		users:  make(map[int]models.User),
		nextId: 1,
	}
}

func (store *InMemoryStore) GetAll() ([]models.User, error) {
	store.Lock()
	defer store.Unlock()

	users := make([]models.User, 0, len(store.users))
	for _, user := range store.users {
		users = append(users, user)
	}

	return users, nil
}

func (store *InMemoryStore) Get(id int) (models.User, error) {
	store.Lock()
	defer store.Unlock()

	user, ok := store.users[id]
	if !ok {
		return models.User{}, fmt.Errorf("user not found (id=%d)", id)
	}

	return user, nil
}

func (store *InMemoryStore) Create(user models.User) (models.User, error) {
	store.Lock()
	defer store.Unlock()

	user.Id = store.nextId

	store.nextId++
	store.users[user.Id] = user

	return user, nil
}
