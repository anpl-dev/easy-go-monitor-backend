package usecase

import (
	"context"
	"go-monitor-tool/internal/domain"
	"time"

	"github.com/google/uuid"
)

type (
	FindUserByEmailUseCase interface {
		Execute(ctx context.Context, input FindUserByEmailInput) (FindUserByEmailOutput, error)
	}

	FindUserByEmailInput struct {
		Email string `json:"email" binding:"required"`
	}

	FindUserByEmailPresenter interface {
		Output(*domain.User) FindUserByEmailOutput
	}

	FindUserByEmailOutput struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	findUserByEmailIntatactor struct {
		repo      domain.UserRepository
		presenter FindUserByEmailPresenter
	}
)

func NewFindUserByEmail(
	repo domain.UserRepository,
	presenter FindUserByEmailPresenter,
) FindUserByEmailUseCase {
	return &findUserByEmailIntatactor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *findUserByEmailIntatactor) Execute(ctx context.Context, input FindUserByEmailInput) (FindUserByEmailOutput, error) {
	user, err := i.repo.FindByEmail(ctx, input.Email)
	if err != nil {
		return FindUserByEmailOutput{}, err
	}
	return i.presenter.Output(user), nil

}
