package usecase

import (
	"context"
	"testing"

	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/errors"

	"github.com/google/uuid"
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
		name          string
		input         DeleteMonitorInput
		repository    domain.MonitorRepository
		expectedError error
	}{
		{
			name: "success: monitor deleted",
			input: DeleteMonitorInput{
				ID: monitor.ID,
			},
			repository:    mockMonitorRepoDelete{err: nil},
			expectedError: nil,
		},
		{
			name: "error: monitor not deleted",
			input: DeleteMonitorInput{
				ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			},
			repository:    mockMonitorRepoDelete{err: errors.ErrNotFound},
			expectedError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewDeleteMonitorInteractor(tt.repository)

			err := uc.Execute(context.Background(), tt.input)
			if tt.expectedError == nil {
				if err != nil {
					t.Errorf("[TestCase '%s'] unexpected error: %v", tt.name, err)
				}
			} else {
				if !errors.Is(err, tt.expectedError) {
					t.Errorf("[TestCase '%s'] Result error: '%v', Expected: '%v'", tt.name, err, tt.expectedError)
				}
			}
		})
	}
}
