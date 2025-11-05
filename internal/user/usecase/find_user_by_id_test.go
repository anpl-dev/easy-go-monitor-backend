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
type mockUserRepoFindByID struct {
	domain.UserRepository

	result *domain.User
	err    error
}

func (m mockUserRepoFindByID) FindByID(_ context.Context, _ uuid.UUID) (*domain.User, error) {
	return m.result, m.err
}

// --- Mock Presenter ---
type mockFindUserByIDPresenter struct {
	result FindUserByIDOutput
}

func (m mockFindUserByIDPresenter) Output(_ *domain.User) FindUserByIDOutput {
	return m.result
}

// --- Test ---
func TestFindUserByIDInteractor_Execute(t *testing.T) {
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
		input         FindUserByIDInput
		mockRepo      mockUserRepoFindByID
		mockPresenter mockFindUserByIDPresenter
		wantError     error
	}{
		{
			name: "success: user found",
			input: FindUserByIDInput{
				ID: user.ID,
			},
			mockRepo: mockUserRepoFindByID{
				result: user,
				err:    nil,
			},
			mockPresenter: mockFindUserByIDPresenter{
				result: FindUserByIDOutput{
					ID:        user.ID,
					Name:      user.Name,
					Email:     user.Email,
					CreatedAt: user.CreatedAt,
					UpdatedAt: user.UpdatedAt,
				},
			},
			wantError: nil,
		},
		{
			name: "error: user not found",
			input: FindUserByIDInput{
				ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			},
			mockRepo: mockUserRepoFindByID{
				result: nil,
				err:    codes.ErrNotFound,
			},
			mockPresenter: mockFindUserByIDPresenter{},
			wantError:     codes.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewFindUserByIDInteractor(&tt.mockRepo, &tt.mockPresenter)
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
