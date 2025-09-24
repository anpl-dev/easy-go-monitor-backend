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

func toDomainUser(s sqlcgen.User) *domain.User {
	return &domain.User{
		ID:           s.ID,
		Name:         s.Name,
		Email:        s.Email,
		PasswordHash: s.PasswordHash,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
	}
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
	return toDomainUser(row), nil
}

func (r *UserPostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	row, err := r.queries.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toDomainUser(row), nil
}

func (r *UserPostgresRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	row, err := r.queries.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return toDomainUser(row), nil
}

func (r *UserPostgresRepository) Update(ctx context.Context, u domain.User) (*domain.User, error) {
	row, err := r.queries.UpdateUser(ctx, sqlcgen.UpdateUserParams{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	})
	if err != nil {
		return nil, err
	}
	return toDomainUser(row), nil
}

func (r *UserPostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.queries.DeleteUser(ctx, id)
}
