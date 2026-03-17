package users

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (User, error) {
	user, err := s.repo.Create(ctx, req)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *Service) GetAllUsers(ctx context.Context) ([]User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) GetUser(ctx context.Context, id int64) (User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *Service) UpdateUser(ctx context.Context, id int64, req UpdateUserRequest) (User, error) {
	user, err := s.repo.Update(ctx, id, req)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *Service) DeleteUser(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
