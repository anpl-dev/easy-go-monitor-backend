package usecase

import (
	"context"
	"testing"

	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/user/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// --- Mock Repository ---
type mockUserRepoDelete struct {
	domain.UserRepository

	err error
}

func (m mockUserRepoDelete) Delete(_ context.Context, _ uuid.UUID) error {
	return m.err
}

// --- Test ---
func TestDeleteUserInteractor_Execute(t *testing.T) {
	t.Parallel()

	user := &domain.User{
		ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
	}

	tests := []struct {
		name      string
		input     DeleteUserInput
		mockRepo  mockUserRepoDelete
		wantError error
	}{
		{
			name: "success: user deleted",
			input: DeleteUserInput{
				ID: user.ID,
			},
			mockRepo:  mockUserRepoDelete{err: nil},
			wantError: nil,
		},
		{
			name: "error: user not deleted",
			input: DeleteUserInput{
				ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			},
			mockRepo:  mockUserRepoDelete{err: codes.ErrNotFound},
			wantError: codes.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewDeleteUserInteractor(&tt.mockRepo)
			err := uc.Execute(context.Background(), tt.input)

			if tt.wantError != nil {
				require.ErrorIs(t, err, tt.wantError, "[%s] error mismatch", tt.name)
			} else {
				require.NoError(t, err, "[%s] unexpected error", tt.name)
			}
		})
	}
}
