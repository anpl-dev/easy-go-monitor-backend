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
		input      CreateMonitorInput
		repository domain.MonitorRepository
		presenter  CreateMonitorPresenter
		want       CreateMonitorOutput
		wantError  error
	}{
		{
			name: "success: create monitor",
			input: CreateMonitorInput{
				UserID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				Name:           "test-monitor",
				URL:            "https://example.com",
				IntervalSecond: 60,
			},
			repository: mockMonitorRepoCreate{
				result: monitor,
				err:    nil,
			},
			presenter: mockCreateMonitorPresenter{
				result: CreateMonitorOutput{
					ID:             monitor.ID,
					UserID:         monitor.UserID,
					Name:           monitor.Name,
					URL:            monitor.URL,
					IntervalSecond: monitor.IntervalSecond,
					CreatedAt:      monitor.CreatedAt,
					UpdatedAt:      monitor.UpdatedAt,
				},
			},
			want: CreateMonitorOutput{
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
			name: "error: missing user id",
			input: CreateMonitorInput{
				Name:           "test-monitor",
				URL:            "https://example.com",
				IntervalSecond: 60,
			},
			repository: mockMonitorRepoCreate{},
			presenter: mockCreateMonitorPresenter{
				result: CreateMonitorOutput{},
			},
			want:      CreateMonitorOutput{},
			wantError: errors.ErrInvalidUUID,
		},
		{
			name: "error: user not found",
			input: CreateMonitorInput{
				UserID:         uuid.MustParse("22222222-2222-2222-2222-222222222222"),
				Name:           "test-monitor",
				URL:            "https://example.com",
				IntervalSecond: 60,
			},
			repository: mockMonitorRepoCreate{
				result: nil,
				err:    errors.ErrNotFound,
			},
			presenter: mockCreateMonitorPresenter{
				result: CreateMonitorOutput{},
			},
			want:      CreateMonitorOutput{},
			wantError: errors.ErrNotFound,
		},
		{
			name: "error: missing name",
			input: CreateMonitorInput{
				UserID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				URL:            "https://example.com",
				IntervalSecond: 60,
			},
			repository: mockMonitorRepoCreate{},
			presenter: mockCreateMonitorPresenter{
				result: CreateMonitorOutput{},
			},
			want:      CreateMonitorOutput{},
			wantError: errors.ErrInvalidMonitorName,
		},
		{
			name: "error: missing url",
			input: CreateMonitorInput{
				UserID:         uuid.MustParse("22222222-2222-2222-2222-222222222222"),
				Name:           "test-monitor",
				IntervalSecond: 60,
			},
			repository: mockMonitorRepoCreate{},
			presenter: mockCreateMonitorPresenter{
				result: CreateMonitorOutput{},
			},
			want:      CreateMonitorOutput{},
			wantError: errors.ErrInvalidMonitorURL,
		},
		{
			name: "error: missing interval",
			input: CreateMonitorInput{
				UserID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				Name:   "test-monitor",
				URL:    "https://example.com",
			},
			repository: mockMonitorRepoCreate{},
			presenter: mockCreateMonitorPresenter{
				result: CreateMonitorOutput{},
			},
			want:      CreateMonitorOutput{},
			wantError: errors.ErrInvalidMonitorInterval,
		},
		{
			name: "error: invalid interval",
			input: CreateMonitorInput{
				UserID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				Name:           "test-monitor",
				URL:            "https://example.com",
				IntervalSecond: -5,
			},
			repository: mockMonitorRepoCreate{},
			presenter: mockCreateMonitorPresenter{
				result: CreateMonitorOutput{},
			},
			want:      CreateMonitorOutput{},
			wantError: errors.ErrInvalidMonitorInterval,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewCreateMonitorInteractor(tt.repository, tt.presenter)

			got, err := uc.Execute(context.Background(), tt.input)
			if tt.wantError == nil {
				if err != nil {
					t.Errorf("[TestCase '%s'] Unexpected error: '%v'", tt.name, err)
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
