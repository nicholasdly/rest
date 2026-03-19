package users

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	store map[string]User
}

func NewService() *Service {
	return &Service{
		store: make(map[string]User),
	}
}

func (s *Service) GetAll() ([]User, error) {
	users := make([]User, 0, len(s.store))

	for _, user := range s.store {
		users = append(users, user)
	}

	return users, nil
}

func (s *Service) Get(id string) (User, error) {
	user, found := s.store[id]
	if !found {
		return User{}, errors.New("user not found")
	}

	return user, nil
}

func (s *Service) Create(username, email string) (User, error) {
	user := User{
		Id:        uuid.NewString(),
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.store[user.Id] = user
	return user, nil
}

func (s *Service) Update(id, username, email string) (User, error) {
	user, found := s.store[id]
	if !found {
		return User{}, errors.New("user not found")
	}

	user.Username = username
	user.Email = email
	user.UpdatedAt = time.Now()

	s.store[id] = user
	return user, nil
}

func (r *Service) Delete(id string) error {
	_, found := r.store[id]
	if !found {
		return errors.New("user not found")
	}

	delete(r.store, id)
	return nil
}
