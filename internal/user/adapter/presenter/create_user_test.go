package presenter

import (
	"easy-go-monitor/internal/user/domain"
	"easy-go-monitor/internal/user/usecase"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserPresenter_Output(t *testing.T) {

	now := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)
	user := domain.User{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		Email:     "alice@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}
	tests := []struct {
		name string
		args *domain.User
		want usecase.CreateUserOutput
	}{
		{
			name: "success: create user",
			args: &user,
			want: usecase.CreateUserOutput{
				ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				Name:      "Alice",
				Email:     "alice@example.com",
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewCreateUserPresenter()
			got := p.Output(tt.args)
			assert.Equal(t, tt.want, got, "[TestCase '%s']", tt.name)
		})
	}
}
