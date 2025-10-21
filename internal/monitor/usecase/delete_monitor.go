package usecase

import (
	"context"
	"easy-go-monitor/internal/monitor/domain"

	"github.com/google/uuid"
)

type (
	// DeleteMonitorUseCase input port
	DeleteMonitorUseCase interface {
		Execute(ctx context.Context, input DeleteMonitorInput) error
	}

	// DeleteMonitorInput input data
	DeleteMonitorInput struct {
		ID uuid.UUID `json:"-"`
	}

	deleteMonitorInteractor struct {
		repo domain.MonitorRepository
	}
)

func NewDeleteMonitorInteractor(repo domain.MonitorRepository) DeleteMonitorUseCase {
	return &deleteMonitorInteractor{repo: repo}
}

func (i *deleteMonitorInteractor) Execute(ctx context.Context, input DeleteMonitorInput) error {
	return i.repo.Delete(ctx, input.ID)
}
