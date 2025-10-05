package usecase

import (
	"context"
	"testing"
	"time"

	"go-monitor-tool/internal/apperr"
	"go-monitor-tool/internal/monitor/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// --- Mock Repository ---
type mockMonitorRepoSearch struct {
	domain.MonitorRepository

	result []*domain.Monitor
	err    error
}

func (m mockMonitorRepoSearch) FindByUserID(_ context.Context, _ uuid.UUID) ([]*domain.Monitor, error) {
	return m.result, m.err
}

// --- Mock Presenter ---
type mockSearchMonitorsPresenter struct {
	result []SearchMonitorsOutput
}

func (m mockSearchMonitorsPresenter) Output(_ []*domain.Monitor) []SearchMonitorsOutput {
	return m.result
}

// --- Test ---
func TestSearchMonitorsInteractor_Execute(t *testing.T) {
	t.Parallel()

	now := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	monitor := &domain.Monitor{
		ID:             uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:           "test-monitor",
		URL:            "https://examaple.com",
		IntervalSecond: 60,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	tests := []struct {
		name          string
		input         SearchMonitorsInput
		mockRepo      mockMonitorRepoSearch
		mockPresenter mockSearchMonitorsPresenter
		wantError     error
	}{
		{
			name: "success: monitors found",
			input: SearchMonitorsInput{
				UserID: "11111111-1111-1111-1111-111111111111",
			},
			mockRepo: mockMonitorRepoSearch{
				result: []*domain.Monitor{monitor},
				err:    nil,
			},
			mockPresenter: mockSearchMonitorsPresenter{
				result: []SearchMonitorsOutput{
					{
						ID:             monitor.ID,
						UserID:         monitor.UserID,
						Name:           monitor.Name,
						URL:            monitor.URL,
						IntervalSecond: monitor.IntervalSecond,
						CreatedAt:      monitor.CreatedAt,
						UpdatedAt:      monitor.UpdatedAt,
					},
				},
			},
			wantError: nil,
		},
		{
			name: "error: monitor not found",
			input: SearchMonitorsInput{
				UserID: "11111111-1111-1111-1111-111111111111",
			},
			mockRepo: mockMonitorRepoSearch{
				result: nil,
				err:    apperr.ErrNotFound,
			},
			mockPresenter: mockSearchMonitorsPresenter{},
			wantError:     apperr.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewSearchMonitorsInteractor(&tt.mockRepo, &tt.mockPresenter)
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
