package presenter

import (
	"easy-go-monitor/internal/user/domain"
	"easy-go-monitor/internal/user/usecase"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUserPresenter_Output(t *testing.T) {

	now := time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)

	user := domain.User{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		Email:     "alice@example.com",
		Password:  "hashedPass",
		CreatedAt: now,
		UpdatedAt: now,
	}

	want := usecase.UpdateUserOutput{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		Email:     "alice@example.com",
		UpdatedAt: now,
	}

	tests := []struct {
		name string
		args *domain.User
		want usecase.UpdateUserOutput
	}{
		{
			name: "success: update user",
			args: &user,
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewUpdateUserPresenter()
			got := p.Output(tt.args)
			assert.Equal(t, tt.want, got, "[TestCase '%s']", tt.name)
		})
	}
}
