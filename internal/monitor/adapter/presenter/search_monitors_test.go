package presenter

import (
	"easy-go-monitor/internal/monitor/domain"
	"easy-go-monitor/internal/monitor/usecase"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSearchMonitorsPresenter_Output(t *testing.T) {

	now := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)

	monitor := domain.Monitor{
		ID:             uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:           "test-monitor",
		URL:            "https://example.com",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	want := usecase.SearchMonitorsOutput{
		ID:             uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:           "test-monitor",
		URL:            "https://example.com",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	tests := []struct {
		name string
		args []*domain.Monitor
		want []usecase.SearchMonitorsOutput
	}{
		{
			name: "success: search monitors",
			args: []*domain.Monitor{&monitor},
			want: []usecase.SearchMonitorsOutput{want},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewSearchMonitorsPresenter()
			got := p.Output(tt.args)
			assert.Equal(t, tt.want, got, "[TestCase '%s']", tt.name)
		})
	}
}
