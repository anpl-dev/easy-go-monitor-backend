package usecase

import (
	"context"
	"testing"

	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/errors"

	"github.com/google/uuid"
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
		name          string
		input         DeleteUserInput
		repository    domain.UserRepository
		wantError error
	}{
		{
			name: "success: user deleted",
			input: DeleteUserInput{
				ID: user.ID,
			},
			repository:    mockUserRepoDelete{err: nil},
			wantError: nil,
		},
		{
			name: "error: user not deleted",
			input: DeleteUserInput{
				ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			},
			repository:    mockUserRepoDelete{err: errors.ErrNotFound},
			wantError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewDeleteUserInteractor(tt.repository)

			err := uc.Execute(context.Background(), tt.input)
			if tt.wantError == nil {
				if err != nil {
					t.Errorf("[TestCase '%s'] Unexpected error: %v", tt.name, err)
				}
			} else {
				if !errors.Is(err, tt.wantError) {
					t.Errorf("[TestCase '%s'] Got error: '%v', Want: '%v'", tt.name, err, tt.wantError)
				}
			}
		})
	}
}
