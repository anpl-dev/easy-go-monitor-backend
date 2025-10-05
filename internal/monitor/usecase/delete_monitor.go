package usecase

import (
	"context"
	"go-monitor-tool/internal/monitor/domain"

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

	deleteMonitorIteractor struct {
		repo domain.MonitorRepository
	}
)

func NewDeleteMonitorInteractor(repo domain.MonitorRepository) DeleteMonitorUseCase {
	return &deleteMonitorIteractor{repo: repo}
}

func (i *deleteMonitorIteractor) Execute(ctx context.Context, input DeleteMonitorInput) error {
	return i.repo.Delete(ctx, input.ID)
}
