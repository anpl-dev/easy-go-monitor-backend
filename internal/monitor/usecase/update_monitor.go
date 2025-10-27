package usecase

import (
	"context"
	"easy-go-monitor/internal/monitor/domain"
	"time"

	"github.com/google/uuid"
)

type (
	// UpdateMonitorUseCase input port
	UpdateMonitorUseCase interface {
		Execute(ctx context.Context, input UpdateMonitorInput) (UpdateMonitorOutput, error)
	}

	// UpdateMonitorInput input data
	UpdateMonitorInput struct {
		ID             uuid.UUID `json:"-"`
		Name           string    `json:"name" binding:"required"`
		URL            string    `json:"url" binding:"required"`
		IntervalSecond int       `json:"interval_second" binding:"required,min=1"`
	}

	// UpdateMonitorPresenter output port
	UpdateMonitorPresenter interface {
		Output(*domain.Monitor) UpdateMonitorOutput
	}

	// UpdateMonitorInput output data
	UpdateMonitorOutput struct {
		ID             uuid.UUID `json:"id"`
		UserID         uuid.UUID `json:"user_id"`
		Name           string    `json:"name"`
		URL            string    `json:"url"`
		IntervalSecond int       `json:"interval_second"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	updateMonitorInteractor struct {
		repo      domain.MonitorRepository
		presenter UpdateMonitorPresenter
	}
)

func NewUpdateMonitorInteractor(
	repo domain.MonitorRepository,
	presenter UpdateMonitorPresenter,
) UpdateMonitorUseCase {
	return &updateMonitorInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *updateMonitorInteractor) Execute(ctx context.Context, input UpdateMonitorInput) (UpdateMonitorOutput, error) {
	monitor := domain.Monitor{
		ID:             input.ID,
		Name:           input.Name,
		URL:            input.URL,
		UpdatedAt:      time.Now(),
	}
	updated, err := i.repo.Update(ctx, monitor)
	if err != nil {
		return UpdateMonitorOutput{}, err
	}

	return i.presenter.Output(updated), nil
}
