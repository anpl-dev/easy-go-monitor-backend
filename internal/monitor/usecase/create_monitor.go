package usecase

import (
	"context"
	"go-monitor-tool/internal/monitor/domain"
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
		UserID         uuid.UUID `json:"user_id"`
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
		ID             uuid.UUID `json:"id"`
		UserID         uuid.UUID `json:"user_id"`
		Name           string    `json:"name"`
		URL            string    `json:"url"`
		IntervalSecond int       `json:"interval_second"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
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
		input.IntervalSecond,
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
