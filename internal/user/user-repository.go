package user

import (
	"context"
	"fmt"

	"github.com/M-kos/crm-user/internal/db"
)

type UserRepository struct {
	db *db.PostgressDB
}

func NewUserRepository(db *db.PostgressDB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetById(ctx context.Context, id string) (*User, error) {
	op := "userRepository.GetById"
	conn, err := r.db.Pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to acquire connection: %w", op, err)
	}

	defer conn.Release()

	row := conn.QueryRow(ctx, "SELECT * FROM users WHERE id = $1", id)

	return nil, nil
}

func (r *UserRepository) Create(ctx context.Context, user *User) error {
	// TODO: Implement the logic to create a new user in the database
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *User) error {
	// TODO: Implement the logic to update an existing user in the database
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	// TODO: Implement the logic to delete a user from the database
	return nil
}

func (r *UserRepository) Get(ctx context.Context) ([]*User, error) {
	// TODO: Implement the logic to retrieve a list of users from the database
	return nil, nil
}
