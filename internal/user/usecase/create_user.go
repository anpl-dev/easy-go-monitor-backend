package usecase

import (
	"context"
	"easy-go-monitor/internal/user/domain"
	"time"

	"github.com/google/uuid"
)

type (
	// CreateUserUseCase input port
	CreateUserUseCase interface {
		Execute(ctx context.Context, input CreateUserInput) (CreateUserOutput, error)
	}

	// CreateUserInput input data
	CreateUserInput struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	// CreateUserPresenter output port
	CreateUserPresenter interface {
		Output(*domain.User) CreateUserOutput
	}

	// CreateUserOutput output data
	CreateUserOutput struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	createUserInteractor struct {
		repo      domain.UserRepository
		presenter CreateUserPresenter
	}
)

func NewCreateUserInteractor(repo domain.UserRepository, presenter CreateUserPresenter) CreateUserUseCase {
	return &createUserInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *createUserInteractor) Execute(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {
	user, err := domain.NewUser(
		input.Name,
		input.Email,
		input.Password,
	)
	if err != nil {
		return CreateUserOutput{}, err
	}

	created, err := i.repo.Create(ctx, *user)
	if err != nil {
		return CreateUserOutput{}, err
	}
	return i.presenter.Output(created), nil
}
