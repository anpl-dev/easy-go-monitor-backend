package usecase

import (
	"context"
	"easy-go-monitor/internal/user/domain"
	"time"

	"github.com/google/uuid"
)

type (
	// FindUserByIDUseCase input port
	FindUserByIDUseCase interface {
		Execute(ctx context.Context, input FindUserByIDInput) (FindUserByIDOutput, error)
	}

	// FindUserByIDInput input data
	FindUserByIDInput struct {
		ID uuid.UUID `json:"id" binding:"required"`
	}

	// FindUserByIDPresenter output port
	FindUserByIDPresenter interface {
		Output(*domain.User) FindUserByIDOutput
	}

	// FindUserByIDOutput output data
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

func NewFindUserByIDInteractor(repo domain.UserRepository, presenter FindUserByIDPresenter) FindUserByIDUseCase {
	return &findUserByIDIntatactor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *findUserByIDIntatactor) Execute(ctx context.Context, input FindUserByIDInput) (FindUserByIDOutput, error) {
	user, err := i.repo.FindByID(ctx, input.ID)
	if err != nil {
		return FindUserByIDOutput{}, err
	}
	return i.presenter.Output(user), nil

}
