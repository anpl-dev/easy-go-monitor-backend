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
type mockUserRepoCreate struct {
	domain.UserRepository

	result *domain.User
	err    error
}

func (m mockUserRepoCreate) Create(_ context.Context, _ domain.User) (*domain.User, error) {
	return m.result, m.err
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

	now := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	hashed, _ := domain.HashedPassword("plainPassword")

	tests := []struct {
		name          string
		input         CreateUserInput
		mockRepo      mockUserRepoCreate
		mockPresenter mockCreateUserPresenter
		wantError     error
	}{
		{
			name: "success: create user",
			input: CreateUserInput{
				Name:     "Alice",
				Email:    "alice@example.com",
				Password: "plainPassword",
			},
			mockRepo: mockUserRepoCreate{
				result: &domain.User{
					ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Name:      "Alice",
					Email:     "alice@example.com",
					Password:  hashed,
					CreatedAt: now,
					UpdatedAt: now,
				},
				err: nil,
			},
			mockPresenter: mockCreateUserPresenter{
				result: CreateUserOutput{
					ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Name:      "Alice",
					Email:     "alice@example.com",
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			wantError: nil,
		},
		{
			name: "error: missing name",
			input: CreateUserInput{
				Name:     "",
				Email:    "alice@exampel.com",
				Password: "plainPassword",
			},
			mockRepo: mockUserRepoCreate{},
			mockPresenter: mockCreateUserPresenter{
				result: CreateUserOutput{},
			},
			wantError: codes.ErrInvalidUserName,
		},
		{
			name: "error: missing email",
			input: CreateUserInput{
				Name:     "Alice",
				Email:    "",
				Password: "plainPassword",
			},
			mockRepo: mockUserRepoCreate{},
			mockPresenter: mockCreateUserPresenter{
				result: CreateUserOutput{},
			},
			wantError: codes.ErrInvalidEmail,
		},
		{
			name: "error: missing password",
			input: CreateUserInput{
				Name:     "Alice",
				Email:    "alice@example.com",
				Password: "",
			},
			mockRepo: mockUserRepoCreate{},
			mockPresenter: mockCreateUserPresenter{
				result: CreateUserOutput{},
			},
			wantError: codes.ErrInvalidPassword,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewCreateUserInteractor(&tt.mockRepo, &tt.mockPresenter)
			got, err := uc.Execute(context.Background(), tt.input)

			if tt.wantError != nil {
				require.ErrorIs(t, err, tt.wantError, "[%s] error mismatch", tt.name)
			} else {
				require.NoError(t, err, "[%s] unexpected error", tt.name)
				require.Equal(t, tt.mockPresenter.result, got, "[%s] output mismatch", tt.name)

				// check hashed password
				require.True(t, domain.CheckPassword(tt.input.Password, tt.mockRepo.result.Password),
					"[%s] password hash mismatch", tt.name)
			}
		})
	}
}
