package usecase

import (
	"context"
	"go-monitor-tool/internal/domain"
	"time"

	"github.com/google/uuid"
)

type (
	SearchUserUseCase interface {
		Execute(ctx context.Context, input SearchUserInput) ([]SearchUserOutput, error)
	}

	SearchUserInput struct {
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	}

	SearchUserPresenter interface {
		Output([]*domain.User) []SearchUserOutput
	}

	SearchUserOutput struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	findUserByEmailIntatactor struct {
		repo      domain.UserRepository
		presenter SearchUserPresenter
	}
)

func NewSearchUserInteractor(repo domain.UserRepository, presenter SearchUserPresenter) SearchUserUseCase {
	return &findUserByEmailIntatactor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *findUserByEmailIntatactor) Execute(ctx context.Context, input SearchUserInput) ([]SearchUserOutput, error) {
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
