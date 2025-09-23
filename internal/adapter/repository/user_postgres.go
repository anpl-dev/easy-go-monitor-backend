package repository

import (
	"context"
	"database/sql"
	"go-monitor-tool/internal/adapter/repository/sqlcgen"
	"go-monitor-tool/internal/domain"

	"github.com/google/uuid"
)

type UserPostgresRepository struct {
	queries *sqlcgen.Queries
}

func NewUserPostgresRepository(db *sql.DB) *UserPostgresRepository {
	return &UserPostgresRepository{queries: sqlcgen.New(db)}
}

func (r *UserPostgresRepository) Create(ctx context.Context, u domain.User) (*domain.User, error) {
	row, err := r.queries.CreateUser(ctx, sqlcgen.CreateUserParams{
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

func (r *UserPostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	u, err := r.queries.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}, nil
}

func (r *UserPostgresRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	u, err := r.queries.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}, nil
}
