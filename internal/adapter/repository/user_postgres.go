package repository

import (
	"context"
	"database/sql"
	"go-monitor-tool/internal/adapter/repository/sqlcgen"
	"go-monitor-tool/internal/domain"
)

type UserPostgresRepository struct {
	q *sqlcgen.Queries
}

func NewUserPostgresRepository(db *sql.DB) *UserPostgresRepository {
	return &UserPostgresRepository{q: sqlcgen.New(db)}
}

func (r *UserPostgresRepository) Create(ctx context.Context, u domain.User) (*domain.User, error) {
	row, err := r.q.CreateUser(ctx, sqlcgen.CreateUserParams{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	})
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:           row.ID,
		Name:         row.Name,
		Email:        row.Email,
		PasswordHash: row.PasswordHash,
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
	}, nil
}
