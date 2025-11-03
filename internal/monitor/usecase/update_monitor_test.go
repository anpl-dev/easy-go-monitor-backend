package usecase

import (
	"context"
	"testing"
	"time"

	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/monitor/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// --- Mock Repository ---
type mockMonitorRepoUpdate struct {
	domain.MonitorRepository

	result *domain.Monitor
	err    error
}

func (m mockMonitorRepoUpdate) Update(_ context.Context, _ domain.Monitor) (*domain.Monitor, error) {
	return m.result, m.err
}

// --- Mock Presenter ---
type mockUpdateMonitorPresenter struct {
	result UpdateMonitorOutput
}

func (m mockUpdateMonitorPresenter) Output(_ *domain.Monitor) UpdateMonitorOutput {
	return m.result
}

// --- Test ---
func TestUpdateMonitorInteractor_Execute(t *testing.T) {
	t.Parallel()

	now := time.Now()
	monitor := &domain.Monitor{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:    uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "test-monitor",
		URL:       "https://examaple.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	tests := []struct {
		name          string
		input         UpdateMonitorInput
		mockRepo      mockMonitorRepoUpdate
		mockPresenter mockUpdateMonitorPresenter
		wantError     error
	}{
		{
			name: "success: monitor updated",
			input: UpdateMonitorInput{
				ID: monitor.ID,
			},
			mockRepo: mockMonitorRepoUpdate{
				result: monitor,
				err:    nil,
			},
			mockPresenter: mockUpdateMonitorPresenter{
				result: UpdateMonitorOutput{
					ID:        monitor.ID,
					UserID:    monitor.UserID,
					Name:      monitor.Name,
					URL:       monitor.URL,
					UpdatedAt: monitor.UpdatedAt,
				},
			},
			wantError: nil,
		},
		{
			name: "error: monitor not updated",
			input: UpdateMonitorInput{
				ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			},
			mockRepo: mockMonitorRepoUpdate{
				result: nil,
				err:    codes.ErrNotFound,
			},
			mockPresenter: mockUpdateMonitorPresenter{},
			wantError:     codes.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewUpdateMonitorInteractor(&tt.mockRepo, &tt.mockPresenter)
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
