package usecase

import (
	"context"
	"easy-go-monitor/internal/infra/logger"
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
		UserID string `json:"-"`
		Name   string `json:"name" binding:"required"`
		URL    string `json:"url" binding:"required"`
		Type   string `json:"type" binding:"required"`
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
		logger    *logger.Logger
	}
)

func NewCreateMonitorInteractor(
	repo domain.MonitorRepository,
	presenter CreateMonitorPresenter,
	logger *logger.Logger,
) CreateMonitorUseCase {
	return &createMonitorInteractor{
		repo:      repo,
		presenter: presenter,
		logger:    logger,
	}
}

func (i *createMonitorInteractor) Execute(ctx context.Context, input CreateMonitorInput) (CreateMonitorOutput, error) {
	i.logger.Debug("CreateMonitor started",
		"user_id", input.UserID,
		"name", input.Name,
		"url", input.URL,
		"type", input.Type,
	)

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return CreateMonitorOutput{}, err
	}

	monitor, err := domain.NewMonitor(
		userID,
		input.Name,
		input.URL,
		input.Type,
	)
	if err != nil {
		i.logger.Error("CreateMonitor", "failed", "error", err)
		return CreateMonitorOutput{}, err
	}

	created, err := i.repo.Create(ctx, *monitor)
	if err != nil {
		return CreateMonitorOutput{}, err
	}

	i.logger.Debug("CreateMonitor success", "monitor_id", monitor.ID.String())
	return i.presenter.Output(created), nil
}
