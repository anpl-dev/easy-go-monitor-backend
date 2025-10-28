package usecase

import (
	"context"
	"easy-go-monitor/internal/monitor/domain"
	"time"

	"github.com/google/uuid"
)

type (
	// CreateMonitorUseCase input port
	CreateMonitorUseCase interface {
		Execute(ctx context.Context, input CreateMonitorInput) (CreateMonitorOutput, error)
	}

	// CreateMonitorInput input data
	CreateMonitorInput struct {
		UserID   uuid.UUID               `json:"user_id"`
		Name     string                  `json:"name" binding:"required"`
		URL      string                  `json:"url" binding:"required"`
		Type     string                  `json:"type" binding:"required"`
		Settings *domain.MonitorSettings `json:"settings" binding:"required"`
	}

	// CreateMonitorPresenter output port
	CreateMonitorPresenter interface {
		Output(*domain.Monitor) CreateMonitorOutput
	}

	// CreateMonitorOutput output data
	CreateMonitorOutput struct {
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

	createMonitorInteractor struct {
		repo      domain.MonitorRepository
		presenter CreateMonitorPresenter
	}
)

func NewCreateMonitorInteractor(
	repo domain.MonitorRepository,
	presenter CreateMonitorPresenter,
) CreateMonitorUseCase {
	return &createMonitorInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *createMonitorInteractor) Execute(ctx context.Context, input CreateMonitorInput) (CreateMonitorOutput, error) {
	monitor, err := domain.NewMonitor(
		input.UserID,
		input.Name,
		input.URL,
		input.Type,
		input.Settings,
	)
	if err != nil {
		return CreateMonitorOutput{}, err
	}

	created, err := i.repo.Create(ctx, *monitor)
	if err != nil {
		return CreateMonitorOutput{}, err
	}

	return i.presenter.Output(created), nil
}
