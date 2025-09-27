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
type mockUserRepoSearch struct {
	domain.UserRepository

	result *domain.User
	err    error
}

func (m mockUserRepoSearch) FindByEmail(_ context.Context, _ string) (*domain.User, error) {
	return m.result, m.err
}

// --- Mock Presenter ---
type mockSearchUsersPresenter struct {
	result []SearchUsersOutput
}

func (m mockSearchUsersPresenter) Output(_ []*domain.User) []SearchUsersOutput {
	return m.result
}

// --- Test ---
func TestSearchUsersInteractor_Execute(t *testing.T) {
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
		input         SearchUsersInput
		repository    domain.UserRepository
		presenter     SearchUsersPresenter
		expected      []SearchUsersOutput
		expectedError error
	}{
		{
			name: "success: user found",
			input: SearchUsersInput{
				Email: user.Email,
				Name:  user.Name,
			},
			repository: mockUserRepoSearch{
				result: user,
				err:    nil,
			},
			presenter: mockSearchUsersPresenter{
				result: []SearchUsersOutput{
					{
						ID:        user.ID,
						Name:      user.Name,
						Email:     user.Email,
						CreatedAt: user.CreatedAt,
						UpdatedAt: user.UpdatedAt,
					},
				},
			},
			expected: []SearchUsersOutput{
				{
					ID:        user.ID,
					Name:      user.Name,
					Email:     user.Email,
					CreatedAt: user.CreatedAt,
					UpdatedAt: user.UpdatedAt,
				},
			},
			expectedError: nil,
		},
		{
			name: "error: user not found",
			input: SearchUsersInput{
				Email: "dummy@example.com",
			},
			repository: mockUserRepoSearch{
				result: nil,
				err:    errors.ErrNotFound,
			},
			presenter:     mockSearchUsersPresenter{},
			expected:      nil,
			expectedError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewSearchUsersInteractor(tt.repository, tt.presenter)

			result, err := uc.Execute(context.Background(), tt.input)
			if tt.expectedError == nil {
				if err != nil {
					t.Errorf("[TestCase '%s'] unexpected error: %v", tt.name, err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("[TestCase '%s'] Result: '%+v' | Expected: '%+v'", tt.name, result, tt.expected)
				}
			} else {
				if !errors.Is(err, tt.expectedError) {
					t.Errorf("[TestCase '%s'] Result error: '%v' | Expected: '%v'", tt.name, err, tt.expectedError)
				}
			}
		})
	}
}
