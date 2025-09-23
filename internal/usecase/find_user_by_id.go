package usecase

import (
	"context"
	"go-monitor-tool/internal/domain"
	"time"

	"github.com/google/uuid"
)

type (
	FindUserByIDUseCase interface {
		Execute(ctx context.Context, input FindUserByIDInput) (FindUserByIDOutput, error)
	}

	FindUserByIDInput struct {
		ID uuid.UUID `json:"id" binding:"required"`
	}

	FindUserByIDPresenter interface {
		Output(*domain.User) FindUserByIDOutput
	}

	FindUserByIDOutput struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	findUserByIDIntatactor struct {
		repo      domain.UserRepository
		presenter FindUserByIDPresenter
	}
)

func NewFindUserByID(
	repo domain.UserRepository,
	presenter FindUserByIDPresenter,
) FindUserByIDUseCase {
	return &findUserByIDIntatactor{
		repo:      repo,
		presenter: presenter,
	}
}

func (uc *findUserByIDIntatactor) Execute(ctx context.Context, input FindUserByIDInput) (FindUserByIDOutput, error) {
	user, err := uc.repo.FindByID(ctx, input.ID)
	if err != nil {
		return FindUserByIDOutput{}, err
	}
	return uc.presenter.Output(user), nil

}
