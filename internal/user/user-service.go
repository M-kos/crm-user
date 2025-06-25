package user

import "context"

type Repository interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*User, error)
	Get(ctx context.Context) ([]*User, error)
}

type UserService struct {
	repo Repository
}

func NewUserService(repo Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, user *User) error {
	return s.repo.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, user *User) error {
	return s.repo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) GetById(ctx context.Context, id string) (*User, error) {
	return s.repo.GetById(ctx, id)
}

func (s *UserService) Get(ctx context.Context) ([]*User, error) {
	return s.repo.Get(ctx)
}
