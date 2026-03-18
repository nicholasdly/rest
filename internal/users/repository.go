package users

import (
	"context"
	"database/sql"
	"log/slog"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll(ctx context.Context) ([]User, error) {
	rows, err := r.db.Query("SELECT id, username, email, created_at FROM users")
	if err != nil {
		slog.Error("failed to query for all users", "err", err)
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt)
		users = append(users, user)
	}

	return users, nil
}

func (r *Repository) Get(ctx context.Context, id int64) (User, error) {
	var user User

	err := r.db.QueryRow(
		"SELECT id, username, email, created_at FROM users WHERE id = ?", id,
	).Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt)

	if err != nil {
		slog.Error("failed to query for specific user", "err", err, "id", id)
		return User{}, err
	}

	return user, nil
}

func (r *Repository) Create(ctx context.Context, req CreateUserRequest) (User, error) {
	var user User

	err := r.db.QueryRow(
		"INSERT INTO users (username, email) VALUES (?, ?) RETURNING id, username, email, created_at", req.Username, req.Email,
	).Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt)

	if err != nil {
		slog.Error("failed to insert new user", "err", err, "req", req)
		return User{}, err
	}

	return user, nil
}

func (r *Repository) Update(ctx context.Context, id int64, req UpdateUserRequest) (User, error) {
	var user User

	err := r.db.QueryRow(
		"UPDATE users SET username = ?, email = ? WHERE id = ? RETURNING id, username, email, created_at", req.Username, req.Email, id,
	).Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt)

	if err != nil {
		slog.Error("failed to update user", "err", err, "req", req)
		return User{}, err
	}

	return user, nil
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)

	if err != nil {
		slog.Error("failed to delete user", "err", err, "id", id)
		return err
	}

	return nil
}
