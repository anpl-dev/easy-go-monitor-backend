package usecase

import (
	"context"
	"go-monitor-tool/internal/user/domain"
	"time"

	"github.com/google/uuid"
)

type (
	// UpdateUserUseCase input port
	UpdateUserUseCase interface {
		Execute(ctx context.Context, input UpdateUserInput) (UpdateUserOutput, error)
	}

	// UpdateUserInput input data
	UpdateUserInput struct {
		ID           uuid.UUID `json:"-"`
		Name         string    `json:"name" binding:"required"`
		Email        string    `json:"email" binding:"required"`
		PasswordHash string    `json:"password_hash" binding:"required"`
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
	updated, err := i.repo.Update(ctx, domain.User{
		ID:           input.ID,
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: input.PasswordHash,
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return UpdateUserOutput{}, err
	}
	return i.presenter.Output(updated), nil
}
