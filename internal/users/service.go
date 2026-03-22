package users

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll(ctx context.Context) ([]User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) Get(ctx context.Context, id string) (User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *Service) Create(ctx context.Context, username, email string) (User, error) {
	user, err := s.repo.Create(ctx, username, email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *Service) Update(ctx context.Context, id, username, email string) (User, error) {
	user, err := s.repo.Update(ctx, id, username, email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
