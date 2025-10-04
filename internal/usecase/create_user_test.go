package usecase

import (
	"context"
	"testing"
	"time"

	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/errors"

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

	now := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	hashed, _ := domain.HashedPassword("plainPassword")

	tests := []struct {
		name      string
		input     CreateUserInput
		mockRepo  mockUserRepoCreate
		presenter mockCreateUserPresenter
		wantError error
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
					ID:           uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Name:         "Alice",
					Email:        "alice@example.com",
					PasswordHash: hashed,
					CreatedAt:    now,
					UpdatedAt:    now,
				},
				err: nil,
			},
			presenter: mockCreateUserPresenter{
				result: CreateUserOutput{
					ID:           uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Name:         "Alice",
					Email:        "alice@example.com",
					CreatedAt:    now,
					UpdatedAt:    now,
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
			presenter: mockCreateUserPresenter{
				result: CreateUserOutput{},
			},
			wantError: errors.ErrInvalidUserName,
		},
		{
			name: "error - missing email",
			input: CreateUserInput{
				Name:     "Alice",
				Email:    "",
				Password: "plainPassword",
			},
			mockRepo: mockUserRepoCreate{},
			presenter: mockCreateUserPresenter{
				result: CreateUserOutput{},
			},
			wantError: errors.ErrInvalidEmail,
		},
		{
			name: "error: missing password",
			input: CreateUserInput{
				Name:     "Alice",
				Email:    "alice@example.com",
				Password: "",
			},
			mockRepo: mockUserRepoCreate{},
			presenter: mockCreateUserPresenter{
				result: CreateUserOutput{},
			},
			wantError: errors.ErrInvalidPassword,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewCreateUserInteractor(&tt.mockRepo, tt.presenter)
			got, err := uc.Execute(context.Background(), tt.input)

			if tt.wantError == nil {
				require.NoError(t, err, "[%s] unexpected error", tt.name)
				require.Equal(t, tt.presenter.result, got, "[%s] output mismatch", tt.name)

				// check hashed password
				require.True(t, domain.CheckPasswordHash(tt.input.Password, tt.mockRepo.result.PasswordHash),
					"[%s] password hash mismatch", tt.name)
			} else {
				require.ErrorIs(t, err, tt.wantError, "[%s] error mismatch", tt.name)
			}
		})
	}
}
