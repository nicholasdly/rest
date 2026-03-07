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

func (s *InMemoryStore) GetAll() ([]models.User, error) {
	s.Lock()
	defer s.Unlock()

	users := make([]models.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}

	return users, nil
}

func (s *InMemoryStore) Get(id int) (models.User, error) {
	s.Lock()
	defer s.Unlock()

	user, found := s.users[id]
	if !found {
		return models.User{}, fmt.Errorf("user not found (id=%d)", id)
	}

	return user, nil
}

func (s *InMemoryStore) Create(user models.User) (models.User, error) {
	s.Lock()
	defer s.Unlock()

	user.Id = s.nextId

	s.nextId++
	s.users[user.Id] = user

	return user, nil
}

func (s *InMemoryStore) Update(user models.User) (models.User, error) {
	s.Lock()
	defer s.Unlock()

	_, found := s.users[user.Id]
	if !found {
		return models.User{}, fmt.Errorf("user not found (id=%d)", user.Id)
	}

	s.users[user.Id] = user
	return user, nil
}

func (s *InMemoryStore) Delete(id int) error {
	s.Lock()
	defer s.Unlock()

	_, found := s.users[id]
	if !found {
		return fmt.Errorf("user not found (id=%d)", id)
	}

	delete(s.users, id)
	return nil
}
