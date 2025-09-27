package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/errors"

	"github.com/google/uuid"
)

// --- Mock Repository ---
type mockCreateUserRepo struct {
	result *domain.User
	err    error
}

func (m mockCreateUserRepo) Create(_ context.Context, _ domain.User) (*domain.User, error) {
	return m.result, m.err
}

// dummy implement
func (m mockCreateUserRepo) FindByID(_ context.Context, _ uuid.UUID) (*domain.User, error) {
	return nil, nil
}

func (m mockCreateUserRepo) FindByEmail(_ context.Context, _ string) (*domain.User, error) {
	return nil, nil
}

func (m mockCreateUserRepo) Update(_ context.Context, _ domain.User) (*domain.User, error) {
	return nil, nil
}

func (m mockCreateUserRepo) Delete(_ context.Context, _ uuid.UUID) error {
	return nil
}

// --- Mock Presenter ---
type mockCreateUserPresenter struct {
	result CreateUserOutput
}

func (m mockCreateUserPresenter) Output(_ *domain.User) CreateUserOutput {
	return m.result
}

// --- Test ---
func TestCreateUserInteractor_Execute(t *testing.T) {
	t.Parallel()

	now := time.Now()
	user := &domain.User{
		ID:           uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:         "Alice",
		Email:        "alice@example.com",
		PasswordHash: "hashedPass",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	tests := []struct {
		name          string
		input         CreateUserInput
		repository    domain.UserRepository
		presenter     CreateUserPresenter
		expected      CreateUserOutput
		expectedError error
	}{
		{
			name: "successs: create user",
			input: CreateUserInput{
				Name:         "Alice",
				Email:        "alice@example.com",
				PasswordHash: "hashedPass",
			},
			repository: mockCreateUserRepo{
				result: user,
				err:    nil,
			},
			presenter: mockCreateUserPresenter{
				result: CreateUserOutput{
					ID:        user.ID,
					Name:      user.Name,
					Email:     user.Email,
					CreatedAt: user.CreatedAt,
					UpdatedAt: user.UpdatedAt,
				},
			},
			expected: CreateUserOutput{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
			expectedError: nil,
		},
		{
			name: "error: missing name",
			input: CreateUserInput{
				Name:         "",
				Email:        "alice@exampel.com",
				PasswordHash: "hashedPass",
			},
			repository: mockCreateUserRepo{},
			presenter: mockCreateUserPresenter{
				result: CreateUserOutput{},
			},
			expectedError: errors.ErrInvalidUserName,
			expected:      CreateUserOutput{},
		},
		{
			name: "error: missing email",
			input: CreateUserInput{
				Name:         "Alice",
				Email:        "",
				PasswordHash: "hashedPass",
			},
			repository: mockCreateUserRepo{},
			presenter: mockCreateUserPresenter{
				result: CreateUserOutput{},
			},
			expectedError: errors.ErrInvalidEmail,
			expected:      CreateUserOutput{},
		},
		{
			name: "error: missing password",
			input: CreateUserInput{
				Name:         "Alice",
				Email:        "alice@example.com",
				PasswordHash: "",
			},
			repository: mockCreateUserRepo{},
			presenter: mockCreateUserPresenter{
				result: CreateUserOutput{},
			},
			expectedError: errors.ErrInvalidPassword,
			expected:      CreateUserOutput{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewCreateUserInteractor(tt.repository, tt.presenter)

			result, err := uc.Execute(context.Background(), tt.input)
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("[TestCase '%s'] Result error: '%v' | Expected: '%v'", tt.name, err, tt.expectedError)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("[TestCase '%s'] Result: '%+v' | Expected: '%+v'", tt.name, result, tt.expected)
			}
		})
	}
}
