package usecase

import (
	"context"
	"testing"
	"time"

	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/infra/logger"
	"easy-go-monitor/internal/monitor/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// --- Mock Repository ---
type mockMonitorRepoCreate struct {
	domain.MonitorRepository

	result *domain.Monitor
	err    error
}

func (m mockMonitorRepoCreate) Create(_ context.Context, _ domain.Monitor) (*domain.Monitor, error) {
	return m.result, m.err
}

// --- Mock Presenter ---
type mockCreateMonitorPresenter struct {
	result CreateMonitorOutput
}

func (m mockCreateMonitorPresenter) Output(_ *domain.Monitor) CreateMonitorOutput {
	return m.result
}

// --- Test ---
func TestCreateMonitorInteractor_Execute(t *testing.T) {
	t.Parallel()

	now := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	monitor := &domain.Monitor{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:    uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "test-monitor",
		URL:       "https://example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	tests := []struct {
		name          string
		input         CreateMonitorInput
		mockRepo      mockMonitorRepoCreate
		mockPresenter mockCreateMonitorPresenter
		want          CreateMonitorOutput
		wantError     error
	}{
		{
			name: "success: create monitor",
			input: CreateMonitorInput{
				Name: "test-monitor",
				URL:  "https://example.com",
			},
			mockRepo: mockMonitorRepoCreate{
				result: monitor,
				err:    nil,
			},
			mockPresenter: mockCreateMonitorPresenter{
				result: CreateMonitorOutput{
					ID:        monitor.ID,
					UserID:    monitor.UserID,
					Name:      monitor.Name,
					URL:       monitor.URL,
					CreatedAt: monitor.CreatedAt,
					UpdatedAt: monitor.UpdatedAt,
				},
			},
			want: CreateMonitorOutput{
				ID:        monitor.ID,
				UserID:    monitor.UserID,
				Name:      monitor.Name,
				URL:       monitor.URL,
				CreatedAt: monitor.CreatedAt,
				UpdatedAt: monitor.UpdatedAt,
			},
			wantError: nil,
		},
		{
			name: "error: missing user id",
			input: CreateMonitorInput{
				Name: "test-monitor",
				URL:  "https://example.com",
			},
			mockRepo: mockMonitorRepoCreate{},
			mockPresenter: mockCreateMonitorPresenter{
				result: CreateMonitorOutput{},
			},
			want:      CreateMonitorOutput{},
			wantError: codes.ErrInvalidUUID,
		},
		{
			name: "error: user not found",
			input: CreateMonitorInput{
				Name: "test-monitor",
				URL:  "https://example.com",
			},
			mockRepo: mockMonitorRepoCreate{
				result: nil,
				err:    codes.ErrNotFound,
			},
			mockPresenter: mockCreateMonitorPresenter{
				result: CreateMonitorOutput{},
			},
			want:      CreateMonitorOutput{},
			wantError: codes.ErrNotFound,
		},
		{
			name: "error: missing name",
			input: CreateMonitorInput{
				URL: "https://example.com",
			},
			mockRepo: mockMonitorRepoCreate{},
			mockPresenter: mockCreateMonitorPresenter{
				result: CreateMonitorOutput{},
			},
			want:      CreateMonitorOutput{},
			wantError: codes.ErrInvalidMonitorName,
		},
		{
			name: "error: missing url",
			input: CreateMonitorInput{
				Name: "test-monitor",
			},
			mockRepo: mockMonitorRepoCreate{},
			mockPresenter: mockCreateMonitorPresenter{
				result: CreateMonitorOutput{},
			},
			want:      CreateMonitorOutput{},
			wantError: codes.ErrInvalidMonitorURL,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewCreateMonitorInteractor(&tt.mockRepo, &tt.mockPresenter, &logger.Logger{})
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
