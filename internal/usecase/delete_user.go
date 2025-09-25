package usecase

import (
	"context"
	"go-monitor-tool/internal/domain"

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

	deleteUserIteractor struct {
		repo domain.UserRepository
	}
)

func NewDeleteUserInteracotr(repo domain.UserRepository) DeleteUserUseCase {
	return &deleteUserIteractor{repo: repo}
}

func (i *deleteUserIteractor) Execute(ctx context.Context, input DeleteUserInput) error {
	return i.repo.Delete(ctx, input.ID)
}
