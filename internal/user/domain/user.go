package domain

import (
	"context"
	"go-monitor-tool/internal/errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

// NewUser creates a new User entity with validation
func NewUser(name, email, plainPassword string) (*User, error) {
	if name == "" {
		return nil, errors.ErrInvalidUserName
	}
	if !strings.Contains(email, "@") {
		return nil, errors.ErrInvalidEmail
	}
	if plainPassword == "" {
		return nil, errors.ErrInvalidPassword
	}

	hashed, err := HashedPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: hashed,
	}, nil
}

func HashedPassword(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (u *User) Authenticate(password string) error {
	if !CheckPasswordHash(password, u.PasswordHash) {
		return errors.ErrInvalidPassword
	}
	return nil
}
