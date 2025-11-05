package usecase

import (
	"context"
	"easy-go-monitor/internal/user/domain"

	"github.com/google/uuid"
)

type (
	// DeleteUserUseCase input port
	DeleteUserUseCase interface {
		Execute(ctx context.Context, input DeleteUserInput) error
	}

	// DeleteUserInput input data
	DeleteUserInput struct {
		ID uuid.UUID `json:"-"`
	}

	deleteUserInteractor struct {
		repo domain.UserRepository
	}
)

func NewDeleteUserInteractor(repo domain.UserRepository) DeleteUserUseCase {
	return &deleteUserInteractor{repo: repo}
}

func (i *deleteUserInteractor) Execute(ctx context.Context, input DeleteUserInput) error {
	return i.repo.Delete(ctx, input.ID)
}
