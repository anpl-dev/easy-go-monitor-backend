package presenter

import (
	"easy-go-monitor/internal/monitor/domain"
	"easy-go-monitor/internal/monitor/usecase"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateMonitorPresenter_Output(t *testing.T) {

	now := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)

	user := domain.Monitor{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:    uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		URL:       "https://example.com",
	}

	want := usecase.UpdateMonitorOutput{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:    uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		URL:       "https://example.com",
		UpdatedAt: now,
	}

	tests := []struct {
		name string
		args *domain.Monitor
		want usecase.UpdateMonitorOutput
	}{
		{
			name: "success: update monitor",
			args: &user,
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewUpdateMonitorPresenter()
			got := p.Output(tt.args)
			assert.Equal(t, tt.want, got, "[TestCase '%s']", tt.name)
		})
	}
}
