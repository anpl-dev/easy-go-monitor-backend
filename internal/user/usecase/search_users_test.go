package usecase

import (
	"context"
	"testing"
	"time"

	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/user/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
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

	now := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	user := &domain.User{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		Email:     "alice@example.com",
		Password:  "hashedPass",
		CreatedAt: now,
		UpdatedAt: now,
	}

	tests := []struct {
		name          string
		input         SearchUsersInput
		mockRepo      mockUserRepoSearch
		mockPresenter mockSearchUsersPresenter
		wantError     error
	}{
		{
			name: "success: user found",
			input: SearchUsersInput{
				Email: user.Email,
				Name:  user.Name,
			},
			mockRepo: mockUserRepoSearch{
				result: user,
				err:    nil,
			},
			mockPresenter: mockSearchUsersPresenter{
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
			wantError: nil,
		},
		{
			name: "error: user not found",
			input: SearchUsersInput{
				Email: "dummy@example.com",
			},
			mockRepo: mockUserRepoSearch{
				result: nil,
				err:    codes.ErrNotFound,
			},
			mockPresenter: mockSearchUsersPresenter{},
			wantError:     codes.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewSearchUsersInteractor(&tt.mockRepo, &tt.mockPresenter)
			got, err := uc.Execute(context.Background(), tt.input)

			if tt.wantError != nil {
				require.ErrorIs(t, err, tt.wantError, "[%s] unexpected err", tt.name)
			} else {
				require.NoError(t, err, "[%s] unexpected err", tt.name)
				require.Equal(t, tt.mockPresenter.result, got, "[%s] result mismatch", tt.name)
			}
		})
	}
}
