package usecase

import (
	"context"
	"go-monitor-tool/internal/domain"
	"time"

	"github.com/google/uuid"
)

type (
	FindMonitorByIDUseCase interface {
		Execute(ctx context.Context, input FindMonitorByIDInput) (FindMonitorByIDOutput, error)
	}

	FindMonitorByIDInput struct {
		ID uuid.UUID `json:"id" binding:"required"`
	}

	FindMonitorByIDPresenter interface {
		Output(*domain.Monitor) FindMonitorByIDOutput
	}

	FindMonitorByIDOutput struct {
		ID             uuid.UUID `json:"id"`
		UserID         uuid.UUID `json:"user_id"`
		Name           string    `json:"name"`
		IntervalSecond int       `json:"interval_second"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	findMonitorByIDInteractor struct {
		repo      domain.MonitorRepository
		presenter FindMonitorByIDPresenter
	}
)

func NewFindMonitorByID(
	repo domain.MonitorRepository,
	presenter FindMonitorByIDPresenter,
) FindMonitorByIDUseCase {
	return &findMonitorByIDInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *findMonitorByIDInteractor) Execute(ctx context.Context, input FindMonitorByIDInput) (FindMonitorByIDOutput, error) {
	monitor, err := i.repo.FindByID(ctx, input.ID)
	if err != nil {
		return FindMonitorByIDOutput{}, err
	}
	if monitor == nil {
		return FindMonitorByIDOutput{}, nil
	}
	return i.presenter.Output(monitor), nil
}
