package usecase

import (
	"context"
	"easy-go-monitor/internal/monitor/domain"
	"time"

	"github.com/google/uuid"
)

type (
	// FindMonitorByIDUseCase input port
	FindMonitorByIDUseCase interface {
		Execute(ctx context.Context, input FindMonitorByIDInput) (FindMonitorByIDOutput, error)
	}

	// FindMonitorByIDInput input data
	FindMonitorByIDInput struct {
		ID uuid.UUID `json:"-"`
	}

	// FindMonitorByIDPresenter output port
	FindMonitorByIDPresenter interface {
		Output(*domain.Monitor) FindMonitorByIDOutput
	}

	// FindMonitorByIDOutput output data
	FindMonitorByIDOutput struct {
		ID        uuid.UUID               `json:"id"`
		UserID    uuid.UUID               `json:"user_id"`
		Name      string                  `json:"name"`
		URL       string                  `json:"url"`
		Type      string                  `json:"type"`
		Settings  *domain.MonitorSettings `json:"settings"`
		IsEnabled bool                    `json:"is_enabled"`
		CreatedAt time.Time               `json:"created_at"`
		UpdatedAt time.Time               `json:"updated_at"`
	}

	findMonitorByIDInteractor struct {
		repo      domain.MonitorRepository
		presenter FindMonitorByIDPresenter
	}
)

func NewFindMonitorByIDInteractor(
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
