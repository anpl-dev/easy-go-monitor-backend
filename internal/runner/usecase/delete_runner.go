package usecase

import (
	"context"
	"easy-go-monitor/internal/runner/domain"

	"github.com/google/uuid"
)

type (
	// DeleteRunnerUseCase input port
	DeleteRunnerUseCase interface {
		Execute(ctx context.Context, input DeleteRunnerInput) error
	}

	// DeleteRunnerInput input data
	DeleteRunnerInput struct {
		ID uuid.UUID `json:"-"`
	}

	deleteRunnerInteractor struct {
		repo domain.RunnerRepository
	}
)

func NewDeleteRunnerInteractor(repo domain.RunnerRepository) DeleteRunnerUseCase {
	return &deleteRunnerInteractor{repo: repo}
}

func (i *deleteRunnerInteractor) Execute(ctx context.Context, input DeleteRunnerInput) error {
	return i.repo.Delete(ctx, input.ID)
}
