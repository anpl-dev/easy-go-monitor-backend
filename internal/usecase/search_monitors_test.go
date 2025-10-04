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
		name       string
		input      SearchMonitorsInput
		repository domain.MonitorRepository
		presenter  SearchMonitorsPresenter
		want       []SearchMonitorsOutput
		wantError  error
	}{
		{
			name: "success: monitors found",
			input: SearchMonitorsInput{
				UserID: "11111111-1111-1111-1111-111111111111",
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
			want: []SearchMonitorsOutput{
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
			wantError: nil,
		},
		{
			name: "error: monitor not found",
			input: SearchMonitorsInput{
				UserID: "11111111-1111-1111-1111-111111111111",
			},
			repository: mockMonitorRepoSearch{
				result: nil,
				err:    errors.ErrNotFound,
			},
			presenter: mockSearchMonitorsPresenter{},
			want:      nil,
			wantError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewSearchMonitorsInteractor(tt.repository, tt.presenter)

			got, err := uc.Execute(context.Background(), tt.input)
			if tt.wantError == nil {
				if err != nil {
					t.Errorf("[TestCase '%s'] unwant error: %v", tt.name, err)
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("[TestCase '%s'] got: '%+v' , want: '%+v'", tt.name, got, tt.want)
				}
			} else {
				if !errors.Is(err, tt.wantError) {
					t.Errorf("[TestCase '%s'] got error: '%v' , want: '%v'", tt.name, err, tt.wantError)
				}
			}
		})
	}
}
