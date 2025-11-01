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
		ID        uuid.UUID               `json:"-"`
		Name      string                  `json:"name" binding:"required"`
		URL       string                  `json:"url" binding:"required"`
		Type      string                  `json:"type" binding:"required"`
		Settings  *domain.MonitorSettings `json:"settings" biding:"required"`
		IsEnabled bool                    `json:"is_enabled" binding:"required"`
	}

	// UpdateMonitorPresenter output port
	UpdateMonitorPresenter interface {
		Output(*domain.Monitor) UpdateMonitorOutput
	}

	// UpdateMonitorInput output data
	UpdateMonitorOutput struct {
		ID        uuid.UUID               `json:"id"`
		UserID    uuid.UUID               `json:"user_id"`
		Name      string                  `json:"name"`
		URL       string                  `json:"url"`
		Type      string                  `json:"type"`
		Settings  *domain.MonitorSettings `json:"settings"`
		IsEnabled bool                    `json:"is_enabled"`
		UpdatedAt time.Time               `json:"updated_at"`
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
		ID:        input.ID,
		Name:      input.Name,
		URL:       input.URL,
		Type:      input.Type,
		Settings:  input.Settings,
		IsEnabled: input.IsEnabled,
	}
	updated, err := i.repo.Update(ctx, monitor)
	if err != nil {
		return UpdateMonitorOutput{}, err
	}

	return i.presenter.Output(updated), nil
}
