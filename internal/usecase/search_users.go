package usecase

import (
	"context"
	"go-monitor-tool/internal/domain"
	"time"

	"github.com/google/uuid"
)

type (
	SearchUsersUseCase interface {
		Execute(ctx context.Context, input SearchUsersInput) ([]SearchUsersOutput, error)
	}

	SearchUsersInput struct {
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	}

	SearchUsersPresenter interface {
		Output([]*domain.User) []SearchUsersOutput
	}

	SearchUsersOutput struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	searchUsersIntatactor struct {
		repo      domain.UserRepository
		presenter SearchUsersPresenter
	}
)

func NewSearchUsersInteractor(repo domain.UserRepository, presenter SearchUsersPresenter) SearchUsersUseCase {
	return &searchUsersIntatactor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *searchUsersIntatactor) Execute(ctx context.Context, input SearchUsersInput) ([]SearchUsersOutput, error) {
	var users []*domain.User
	var err error

	if input.Email != "" {
		user, err := i.repo.FindByEmail(ctx, input.Email)
		if err != nil {
			return nil, err
		}
		users = []*domain.User{user}
	} else if input.Name != "" {
		return nil, nil
	}
	return i.presenter.Output(users), err

}
