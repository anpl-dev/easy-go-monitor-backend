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
type mockUserRepoUpdate struct {
	domain.UserRepository

	result *domain.User
	err    error
}

func (m mockUserRepoUpdate) Update(_ context.Context, _ domain.User) (*domain.User, error) {
	return m.result, m.err
}

// --- Mock Presenter ---
type mockUpdateUserPresenter struct {
	result UpdateUserOutput
}

func (m mockUpdateUserPresenter) Output(_ *domain.User) UpdateUserOutput {
	return m.result
}

// --- Test ---
func TestUpdateUserInteractor_Execute(t *testing.T) {
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
		input         UpdateUserInput
		repository    domain.UserRepository
		presenter     UpdateUserPresenter
		expected      UpdateUserOutput
		expectedError error
	}{
		{
			name: "success: user updated",
			input: UpdateUserInput{
				ID: user.ID,
			},
			repository: mockUserRepoUpdate{
				result: user,
				err:    nil,
			},
			presenter: mockUpdateUserPresenter{
				result: UpdateUserOutput{
					ID:        user.ID,
					Name:      user.Name,
					Email:     user.Email,
					UpdatedAt: user.UpdatedAt,
				},
			},
			expected: UpdateUserOutput{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				UpdatedAt: user.UpdatedAt,
			},
			expectedError: nil,
		},
		{
			name: "error: user not updated",
			input: UpdateUserInput{
				ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			},
			repository: mockUserRepoUpdate{
				result: nil,
				err:    errors.ErrNotFound,
			},
			presenter:     mockUpdateUserPresenter{},
			expected:      UpdateUserOutput{},
			expectedError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewUpdateUserInteractor(tt.repository, tt.presenter)

			result, err := uc.Execute(context.Background(), tt.input)

			if tt.expectedError == nil {
				if err != nil {
					t.Errorf("[TestCase '%s'] Unexpected error: %v", tt.name, err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("[TestCase '%s'] Got: '%+v' , Want: '%+v'", tt.name, result, tt.expected)
				}
			} else {
				if !errors.Is(err, tt.expectedError) {
					t.Errorf("[TestCase '%s'] Got error: '%v' , Want: '%v'", tt.name, err, tt.expectedError)
				}
			}
		})
	}
}
