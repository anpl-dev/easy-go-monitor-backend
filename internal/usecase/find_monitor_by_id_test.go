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
		input         FindMonitorByIDInput
		repository    domain.MonitorRepository
		presenter     FindMonitorByIDPresenter
		want      FindMonitorByIDOutput
		wantError error
	}{
		{
			name: "success: monitor found",
			input: FindMonitorByIDInput{
				ID: monitor.ID,
			},
			repository: mockMonitorRepoFindByID{
				result: monitor,
				err:    nil,
			},
			presenter: mockFindMonitorByIDPresenter{
				result: FindMonitorByIDOutput{
						ID:             monitor.ID,
						UserID:         monitor.UserID,
						Name:           monitor.Name,
						URL:            monitor.URL,
						IntervalSecond: monitor.IntervalSecond,
						CreatedAt:      monitor.CreatedAt,
						UpdatedAt:      monitor.UpdatedAt,
				},
			},
			want: FindMonitorByIDOutput{
						ID:             monitor.ID,
						UserID:         monitor.UserID,
						Name:           monitor.Name,
						URL:            monitor.URL,
						IntervalSecond: monitor.IntervalSecond,
						CreatedAt:      monitor.CreatedAt,
						UpdatedAt:      monitor.UpdatedAt,
			},
			wantError: nil,
		},
		{
			name: "error: monitor not found",
			input: FindMonitorByIDInput{
				ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			},
			repository: mockMonitorRepoFindByID{
				result: nil,
				err:    errors.ErrNotFound,
			},
			presenter:     mockFindMonitorByIDPresenter{},
			want:      FindMonitorByIDOutput{},
			wantError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewFindMonitorByIDInteractor(tt.repository, tt.presenter)

			got, err := uc.Execute(context.Background(), tt.input)
			if tt.wantError == nil {
				if err != nil {
					t.Errorf("[TestCase '%s'] Unexpected error: %v", tt.name, err)
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("[TestCase '%s'] Got: '%+v' , Want: '%+v'", tt.name, got, tt.want)
				}
			} else {
				if !errors.Is(err, tt.wantError) {
					t.Errorf("[TestCase '%s'] Got error: '%v' , Want: '%v'", tt.name, err, tt.wantError)
				}
			}
		})
	}
}
