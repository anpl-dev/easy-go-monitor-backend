package repository

import (
	"context"
	"database/sql"
	"easy-go-monitor/db/sqlcgen"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/user/domain"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPostgresRepository struct {
	queries *sqlcgen.Queries
}

func NewUserPostgresRepository(pool *pgxpool.Pool) *UserPostgresRepository {
	return &UserPostgresRepository{queries: sqlcgen.New(pool)}
}

func toDomainUser(s sqlcgen.User) *domain.User {
	return &domain.User{
		ID:        s.ID,
		Name:      s.Name,
		Email:     s.Email,
		Password:  s.Password,
		CreatedAt: *s.CreatedAt,
		UpdatedAt: *s.UpdatedAt,
	}
}

// Create
func (r *UserPostgresRepository) Create(ctx context.Context, u domain.User) (*domain.User, error) {
	row, err := r.queries.CreateUser(ctx, sqlcgen.CreateUserParams{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		return nil, err
	}
	return toDomainUser(row), nil
}

// FindByID
func (r *UserPostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	row, err := r.queries.FindUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, codes.ErrNotFound
		}
		return nil, err
	}
	return toDomainUser(row), nil
}

// FindByEmail
func (r *UserPostgresRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	row, err := r.queries.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, codes.ErrNotFound
		}
		return nil, err
	}
	return toDomainUser(row), nil
}

// Update
func (r *UserPostgresRepository) Update(ctx context.Context, u domain.User) (*domain.User, error) {
	row, err := r.queries.UpdateUser(ctx, sqlcgen.UpdateUserParams{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, codes.ErrNotFound
		}
		return nil, err
	}
	return toDomainUser(row), nil
}

// Delete
func (r *UserPostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.queries.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return codes.ErrNotFound
		}
		return err
	}
	return nil
}
