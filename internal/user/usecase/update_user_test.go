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
		input         UpdateUserInput
		mockRepo      mockUserRepoUpdate
		mockPresenter mockUpdateUserPresenter
		wantError     error
	}{
		{
			name: "success: user updated",
			input: UpdateUserInput{
				ID: user.ID,
			},
			mockRepo: mockUserRepoUpdate{
				result: user,
				err:    nil,
			},
			mockPresenter: mockUpdateUserPresenter{
				result: UpdateUserOutput{
					ID:        user.ID,
					Name:      user.Name,
					Email:     user.Email,
					UpdatedAt: user.UpdatedAt,
				},
			},
			wantError: nil,
		},
		{
			name: "error: user not updated",
			input: UpdateUserInput{
				ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			},
			mockRepo: mockUserRepoUpdate{
				result: nil,
				err:    codes.ErrNotFound,
			},
			mockPresenter: mockUpdateUserPresenter{},
			wantError:     codes.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewUpdateUserInteractor(&tt.mockRepo, &tt.mockPresenter)
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
