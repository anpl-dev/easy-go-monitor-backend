package usecase

import (
	"context"
	"easy-go-monitor/internal/monitor/domain"
	"time"

	"github.com/google/uuid"
)

type (
	SetEnabledMonitorUseCase interface {
		Execute(ctx context.Context, input SetEnabledMonitorInput) (SetEnabledMonitorOutput, error)
	}

	SetEnabledMonitorInput struct {
		ID        uuid.UUID `json:"-"`
		IsEnabled bool      `json:"is_enabled"`
	}

	SetEnabledMonitorPresenter interface {
		Output(*domain.Monitor) SetEnabledMonitorOutput
	}

	SetEnabledMonitorOutput struct {
		ID        uuid.UUID               `json:"id"`
		UserID    uuid.UUID               `json:"user_id"`
		Name      string                  `json:"name"`
		URL       string                  `json:"url"`
		Type      string                  `json:"type"`
		Settings  *domain.MonitorSettings `json:"settings"`
		IsEnabled bool                    `json:"is_enabled"`
		UpdatedAt time.Time               `json:"updated_at"`
	}

	setEnabledMonitorInteractor struct {
		repo      domain.MonitorRepository
		presenter SetEnabledMonitorPresenter
	}
)

func NewSetEnabledMonitorInteractor(
	repo domain.MonitorRepository,
	presenter SetEnabledMonitorPresenter,
) SetEnabledMonitorUseCase {
	return &setEnabledMonitorInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *setEnabledMonitorInteractor) Execute(ctx context.Context, input SetEnabledMonitorInput) (SetEnabledMonitorOutput, error) {
	monitor, err := i.repo.SetEnabled(ctx, input.ID, input.IsEnabled)
	if err != nil {
		return SetEnabledMonitorOutput{}, err
	}
	return i.presenter.Output(monitor), nil
}
