package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"go-monitor-tool/internal/domain"

	"github.com/google/uuid"
)

// --- Mock Repository ---
type mockUserRepoStore struct {
	result *domain.User
	err    error
}

func (m mockUserRepoStore) Create(_ context.Context, _ domain.User) (*domain.User, error) {
	return m.result, m.err
}

// dummy implement
func (m mockUserRepoStore) FindByID(_ context.Context, _ uuid.UUID) (*domain.User, error) {
	return nil, nil
}

func (m mockUserRepoStore) FindByEmail(_ context.Context, _ string) (*domain.User, error) {
	return nil, nil
}

func (m mockUserRepoStore) Update(_ context.Context, _ domain.User) (*domain.User, error) {
	return nil, nil
}

func (m mockUserRepoStore) Delete(_ context.Context, _ uuid.UUID) error {
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
		ID:           uuid.MustParse("74ce6ef9-d96e-43dd-8be4-3b7f0b5dbef5"),
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
			name: "Create user successful",
			input: CreateUserInput{
				Name:         "Alice",
				Email:        "alice@example.com",
				PasswordHash: "hashedPass",
			},
			repository: mockUserRepoStore{
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
