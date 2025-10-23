package presenter

import (
	"easy-go-monitor/internal/user/domain"
	"easy-go-monitor/internal/user/usecase"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSearchUsersPresenter_Output(t *testing.T) {

	now := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)

	user := domain.User{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		Email:     "alice@example.com",
		Password:  "hashedPass",
		CreatedAt: now,
		UpdatedAt: now,
	}

	want := usecase.SearchUsersOutput{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		Email:     "alice@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	tests := []struct {
		name string
		args []*domain.User
		want []usecase.SearchUsersOutput
	}{
		{
			name: "success: search users",
			args: []*domain.User{&user},
			want: []usecase.SearchUsersOutput{want},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewSearchUsersPresenter()
			got := p.Output(tt.args)
			assert.Equal(t, tt.want, got, "[TestCase '%s']", tt.name)
		})
	}
}
