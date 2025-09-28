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
		input         UpdateMonitorInput
		repository    domain.MonitorRepository
		presenter     UpdateMonitorPresenter
		expected      UpdateMonitorOutput
		expectedError error
	}{
		{
			name: "success: monitor updated",
			input: UpdateMonitorInput{
				ID: monitor.ID,
			},
			repository: mockMonitorRepoUpdate{
				result: monitor,
				err:    nil,
			},
			presenter: mockUpdateMonitorPresenter{
				result: UpdateMonitorOutput{
					ID:             monitor.ID,
					UserID:         monitor.UserID,
					Name:           monitor.Name,
					URL:            monitor.URL,
					IntervalSecond: monitor.IntervalSecond,
					UpdatedAt:      monitor.UpdatedAt,
				},
			},
			expected: UpdateMonitorOutput{
				ID:             monitor.ID,
				UserID:         monitor.UserID,
				Name:           monitor.Name,
				URL:            monitor.URL,
				IntervalSecond: monitor.IntervalSecond,
				UpdatedAt:      monitor.UpdatedAt,
			},
			expectedError: nil,
		},
		{
			name: "error: monitor not updated",
			input: UpdateMonitorInput{
				ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			},
			repository: mockMonitorRepoUpdate{
				result: nil,
				err:    errors.ErrNotFound,
			},
			presenter:     mockUpdateMonitorPresenter{},
			expected:      UpdateMonitorOutput{},
			expectedError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewUpdateMonitorInteractor(tt.repository, tt.presenter)

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
