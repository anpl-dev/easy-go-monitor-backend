package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/errors"

	"github.com/google/uuid"
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

	now := time.Now()
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
		repository    domain.MonitorRepository
		presenter     SearchMonitorsPresenter
		expected      []SearchMonitorsOutput
		expectedError error
	}{
		{
			name: "success: monitors found",
			input: SearchMonitorsInput{
				UserID: monitor.UserID,
			},
			repository: mockMonitorRepoSearch{
				result: []*domain.Monitor{monitor},
				err:    nil,
			},
			presenter: mockSearchMonitorsPresenter{
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
			expected: []SearchMonitorsOutput{
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
			expectedError: nil,
		},
		{
			name:  "error: monitor not found",
			input: SearchMonitorsInput{},
			repository: mockMonitorRepoSearch{
				result: nil,
				err:    errors.ErrNotFound,
			},
			presenter:     mockSearchMonitorsPresenter{},
			expected:      nil,
			expectedError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewSearchMonitorsInteractor(tt.repository, tt.presenter)

			result, err := uc.Execute(context.Background(), tt.input)
			if tt.expectedError == nil {
				if err != nil {
					t.Errorf("[TestCase '%s'] unexpected error: %v", tt.name, err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("[TestCase '%s'] Result: '%+v' | Expected: '%+v'", tt.name, result, tt.expected)
				}
			} else {
				if !errors.Is(err, tt.expectedError) {
					t.Errorf("[TestCase '%s'] Result error: '%v' | Expected: '%v'", tt.name, err, tt.expectedError)
				}
			}
		})
	}
}
