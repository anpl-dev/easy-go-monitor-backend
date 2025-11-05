package usecase

import (
	"context"
	"testing"

	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/monitor/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// --- Mock Repository ---
type mockMonitorRepoDelete struct {
	domain.MonitorRepository

	err error
}

func (m mockMonitorRepoDelete) Delete(_ context.Context, _ uuid.UUID) error {
	return m.err
}

// --- Test ---
func TestDeleteMonitorInteractor_Execute(t *testing.T) {
	t.Parallel()

	monitor := &domain.Monitor{
		ID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
	}

	tests := []struct {
		name      string
		input     DeleteMonitorInput
		mockRepo  mockMonitorRepoDelete
		wantError error
	}{
		{
			name: "success: monitor deleted",
			input: DeleteMonitorInput{
				ID: monitor.ID,
			},
			mockRepo:  mockMonitorRepoDelete{err: nil},
			wantError: nil,
		},
		{
			name: "error: monitor not deleted",
			input: DeleteMonitorInput{
				ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			},
			mockRepo:  mockMonitorRepoDelete{err: codes.ErrNotFound},
			wantError: codes.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewDeleteMonitorInteractor(&tt.mockRepo)
			err := uc.Execute(context.Background(), tt.input)

			if tt.wantError != nil {
				require.ErrorIs(t, err, tt.wantError, "[%s] error mismatch", tt.name)
			} else {
				require.NoError(t, err, "[%s] unexpected error", tt.name)
			}
		})
	}
}
