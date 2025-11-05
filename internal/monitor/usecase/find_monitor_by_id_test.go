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
type mockMonitorRepoFindByID struct {
	domain.MonitorRepository

	result *domain.Monitor
	err    error
}

func (m mockMonitorRepoFindByID) FindByID(_ context.Context, _ uuid.UUID) (*domain.Monitor, error) {
	return m.result, m.err
}

// --- Mock Presenter ---
type mockFindMonitorByIDPresenter struct {
	result FindMonitorByIDOutput
}

func (m mockFindMonitorByIDPresenter) Output(_ *domain.Monitor) FindMonitorByIDOutput {
	return m.result
}

// --- Test ---
func TestFindMonitorByIDInteractor_Execute(t *testing.T) {
	t.Parallel()

	now := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
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
		input         FindMonitorByIDInput
		mockRepo      mockMonitorRepoFindByID
		mockPresenter mockFindMonitorByIDPresenter
		wantError     error
	}{
		{
			name: "success: monitor found",
			input: FindMonitorByIDInput{
				ID: monitor.ID,
			},
			mockRepo: mockMonitorRepoFindByID{
				result: monitor,
				err:    nil,
			},
			mockPresenter: mockFindMonitorByIDPresenter{
				result: FindMonitorByIDOutput{
					ID:        monitor.ID,
					UserID:    monitor.UserID,
					Name:      monitor.Name,
					URL:       monitor.URL,
					CreatedAt: monitor.CreatedAt,
					UpdatedAt: monitor.UpdatedAt,
				},
			},
			wantError: nil,
		},
		{
			name: "error: monitor not found",
			input: FindMonitorByIDInput{
				ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			},
			mockRepo: mockMonitorRepoFindByID{
				result: nil,
				err:    codes.ErrNotFound,
			},
			mockPresenter: mockFindMonitorByIDPresenter{},
			wantError:     codes.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewFindMonitorByIDInteractor(&tt.mockRepo, &tt.mockPresenter)
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
