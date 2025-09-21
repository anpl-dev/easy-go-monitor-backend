package usecase

import (
	"context"
	"go-monitor-tool/internal/domain"
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
		UserID         uuid.UUID `json:"user_id" binding:"required"`
		Name           string    `json:"name" binding:"required"`
		URL            string    `json:"url" binding:"required"`
		IntervalSecond int       `json:"interval_second" binding:"required,min=1"`
	}

	// CreateMonitorPresenter output port
	CreateMonitorPresenter interface {
		Output(*domain.Monitor) CreateMonitorOutput
	}

	// CreateMonitorInput output data
	CreateMonitorOutput struct {
		ID             uuid.UUID
		UserID         uuid.UUID
		Name           string
		URL            string
		IntervalSecond int
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}

	createMonitorInteractor struct {
		repo      domain.MonitorRepository
		presenter CreateMonitorPresenter
	}
)

func NewCreateMonitor(repo domain.MonitorRepository, presenter CreateMonitorPresenter) CreateMonitorUseCase {
	return &createMonitorInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (m *createMonitorInteractor) Execute(ctx context.Context, input CreateMonitorInput) (CreateMonitorOutput, error) {
	monitor, err := domain.NewMonitor(
		input.UserID,
		input.Name,
		input.URL,
		input.IntervalSecond,
	)
	if err != nil {
		return CreateMonitorOutput{}, err
	}

	created, err := m.repo.Create(ctx, *monitor)
	if err != nil {
		return CreateMonitorOutput{}, err
	}

	return m.presenter.Output(created), nil
}
