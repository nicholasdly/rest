package users

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll(ctx context.Context) ([]User, error) {
	query := `SELECT id, username, email, created_at, updated_at FROM users`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) Get(ctx context.Context, id string) (User, error) {
	var user User

	query := `
		SELECT id, username, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *Repository) Create(ctx context.Context, username, email string) (User, error) {
	var user User

	query := `
		INSERT INTO users (username, email)
		VALUES ($1, $2)
		RETURNING id, username, email, created_at, updated_at;
	`

	err := r.db.QueryRow(ctx, query, username, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *Repository) Update(ctx context.Context, id, username, email string) (User, error) {
	var user User

	query := `
		UPDATE users
		SET username = $2, email = $3
		WHERE id = $1
		RETURNING id, username, email, created_at, updated_at;
	`

	err := r.db.QueryRow(ctx, query, id, username, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1;`

	_, err := r.db.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
