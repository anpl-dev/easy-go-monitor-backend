package domain

import (
	"context"
	"go-monitor-tool/internal/errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type (
	UserRepository interface {
		Create(ctx context.Context, u User) (*User, error)
		FindByID(ctx context.Context, id uuid.UUID) (*User, error)
		FindByEmail(ctx context.Context, email string) (*User, error)
		Update(ctx context.Context, u User) (*User, error)
		Delete(ctx context.Context, id uuid.UUID) error
	}

	User struct {
		ID           uuid.UUID
		Name         string
		Email        string
		PasswordHash string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}
)

// validation: User
func NewUser(name, email, passwordHash string) (*User, error) {
	if name == "" {
		return nil, errors.ErrInvalidUserName
	}
	if !strings.Contains(email, "@") {
		return nil, errors.ErrInvalidEmail
	}
	if passwordHash == "" {
		return nil, errors.ErrInvalidPassword
	}

	return &User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
	}, nil
}
