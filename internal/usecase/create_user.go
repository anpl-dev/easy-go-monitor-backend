package usecase

import (
	"context"
	"go-monitor-tool/internal/domain"
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
		Name         string `json:"name" binding:"required"`
		Email        string `json:"email" binding:"required"`
		PasswordHash string `json:"password_hash" binding:"required"`
	}
	// CreateUserPresenter output port
	CreateUserPresenter interface {
		Output(*domain.User) CreateUserOutput
	}

	// CreateUserOutput output data
	CreateUserOutput struct {
		ID           uuid.UUID `json:"id"`
		Name         string    `json:"name"`
		Email        string    `json:"email"`
		PasswordHash string    `json:"password_hash"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	createUserIteractor struct {
		repo      domain.UserRepository
		presenter CreateUserPresenter
	}
)

func NewCreateUserInteractor(
	repo domain.UserRepository,
	presenter CreateUserPresenter,
) CreateUserUseCase {
	return &createUserIteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *createUserIteractor) Execute(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {
	user, err := domain.NewUser(
		input.Name,
		input.Email,
		input.PasswordHash,
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
