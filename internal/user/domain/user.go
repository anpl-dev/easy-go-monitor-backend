package domain

import (
	"context"
	"easy-go-monitor/internal/codes"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserRepository interface {
		Create(ctx context.Context, user User) (*User, error)
		FindByID(ctx context.Context, id uuid.UUID) (*User, error)
		FindByEmail(ctx context.Context, email string) (*User, error)
		// FindAll(ctx context.Context) ([]*User, error)
		Update(ctx context.Context, user User) (*User, error)
		Delete(ctx context.Context, id uuid.UUID) error
	}

	User struct {
		ID        uuid.UUID
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

// NewUser creates a new User entity with validation
func NewUser(name, email, plainPassword string) (*User, error) {
	if name == "" {
		return nil, codes.ErrInvalidUserName
	}
	if !strings.Contains(email, "@") {
		return nil, codes.ErrInvalidEmail
	}
	if plainPassword == "" {
		return nil, codes.ErrInvalidPassword
	}

	hashedPassword, err := HashedPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}, nil
}

func HashedPassword(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (u *User) Authenticate(password string) error {
	if !CheckPassword(password, u.Password) {
		return codes.ErrInvalidPassword
	}
	return nil
}
