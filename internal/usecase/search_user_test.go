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
type mockSearchUserPresenter struct {
	result []SearchUserOutput
}

func (m mockSearchUserPresenter) Output(_ []*domain.User) []SearchUserOutput {
	return m.result
}

// --- Test ---
func TestSearchUserInteractor_Execute(t *testing.T) {
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
		input         SearchUserInput
		repository    domain.UserRepository
		presenter     SearchUserPresenter
		expected      []SearchUserOutput
		expectedError error
	}{
		{
			name: "successs: user found",
			input: SearchUserInput{
				Email: user.Email,
				Name:  user.Name,
			},
			repository: mockUserRepoSearch{
				result: user,
				err:    nil,
			},
			presenter: mockSearchUserPresenter{
				result: []SearchUserOutput{
					{
						ID:        user.ID,
						Name:      user.Name,
						Email:     user.Email,
						CreatedAt: user.CreatedAt,
						UpdatedAt: user.UpdatedAt,
					},
				},
			},
			expected: []SearchUserOutput{
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
			input: SearchUserInput{
				Email: "dummy@example.com",
			},
			repository: mockUserRepoSearch{
				result: nil,
				err:    errors.ErrNotFound,
			},
			presenter:     mockSearchUserPresenter{},
			expected:      nil,
			expectedError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewSearchUserInteractor(tt.repository, tt.presenter)

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
