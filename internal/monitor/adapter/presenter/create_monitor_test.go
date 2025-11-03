package presenter

import (
	"easy-go-monitor/internal/monitor/domain"
	"easy-go-monitor/internal/monitor/usecase"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateMonitorPresenter_Output(t *testing.T) {

	now := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	user := domain.Monitor{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "test-monitor",
		URL:       "https://example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}
	tests := []struct {
		name string
		args *domain.Monitor
		want usecase.CreateMonitorOutput
	}{
		{
			name: "success: create user",
			args: &user,
			want: usecase.CreateMonitorOutput{
				ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				Name:      "test-monitor",
				URL:       "https://example.com",
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewCreateMonitorPresenter()
			got := p.Output(tt.args)
			assert.Equal(t, tt.want, got, "[TestCase '%s']", tt.name)
		})
	}
}
