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
type mockUserRepoFindByID struct {
	domain.UserRepository

	result *domain.User
	err    error
}

func (m mockUserRepoFindByID) FindByID(_ context.Context, _ uuid.UUID) (*domain.User, error) {
	return m.result, m.err
}

// --- Mock Presenter ---
type mockFindUserByIDPresenter struct {
	result FindUserByIDOutput
}

func (m mockFindUserByIDPresenter) Output(_ *domain.User) FindUserByIDOutput {
	return m.result
}

// --- Test ---
func TestFindUserByIDInteractor_Execute(t *testing.T) {
	t.Parallel()

	now := time.Now()
	user := &domain.User{
		ID:           uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:         "Alice",
		Email:        "alice@example.com",
		PasswordHash: "hashedPass",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	tests := []struct {
		name          string
		input         FindUserByIDInput
		repository    domain.UserRepository
		presenter     FindUserByIDPresenter
		want      FindUserByIDOutput
		wantError error
	}{
		{
			name: "success: user found",
			input: FindUserByIDInput{
				ID: user.ID,
			},
			repository: mockUserRepoFindByID{
				result: user,
				err:    nil,
			},
			presenter: mockFindUserByIDPresenter{
				result: FindUserByIDOutput{
					ID:        user.ID,
					Name:      user.Name,
					Email:     user.Email,
					CreatedAt: user.CreatedAt,
					UpdatedAt: user.UpdatedAt,
				},
			},
			want: FindUserByIDOutput{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
			wantError: nil,
		},
		{
			name: "error: user not found",
			input: FindUserByIDInput{
				ID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			},
			repository: mockUserRepoFindByID{
				result: nil,
				err:    errors.ErrNotFound,
			},
			presenter:     mockFindUserByIDPresenter{},
			want:      FindUserByIDOutput{},
			wantError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewFindUserByIDInteractor(tt.repository, tt.presenter)

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
