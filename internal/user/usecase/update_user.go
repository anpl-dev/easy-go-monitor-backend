package usecase

import (
	"context"
	"easy-go-monitor/internal/user/domain"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	// UpdateUserUseCase input port
	UpdateUserUseCase interface {
		Execute(ctx context.Context, input UpdateUserInput) (UpdateUserOutput, error)
	}

	// UpdateUserInput input data
	UpdateUserInput struct {
		ID       uuid.UUID `json:"-"`
		Name     string    `json:"name" binding:"required"`
		Email    string    `json:"email" binding:"required"`
		Password string    `json:"password" binding:"required"`
	}
	// UpdateUserPresenter output port
	UpdateUserPresenter interface {
		Output(*domain.User) UpdateUserOutput
	}

	// UpdateUserOutput output data
	UpdateUserOutput struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	updateUserInteractor struct {
		repo      domain.UserRepository
		presenter UpdateUserPresenter
	}
)

func NewUpdateUserInteractor(
	repo domain.UserRepository,
	presenter UpdateUserPresenter,
) UpdateUserUseCase {
	return &updateUserInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *updateUserInteractor) Execute(ctx context.Context, input UpdateUserInput) (UpdateUserOutput, error) {
	var hashedPassword string
	if input.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return UpdateUserOutput{}, err
		}
		hashedPassword = string(hash)
	}

	updated, err := i.repo.Update(ctx, domain.User{
		ID:        input.ID,
		Name:      input.Name,
		Email:     input.Email,
		Password:  hashedPassword,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return UpdateUserOutput{}, err
	}
	return i.presenter.Output(updated), nil
}
