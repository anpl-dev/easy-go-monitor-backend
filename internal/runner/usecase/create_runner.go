package usecase

import (
	"context"
	"easy-go-monitor/internal/runner/domain"
	"time"

	"github.com/google/uuid"
)

type (
	// CreateRunnerUseCase input port
	CreateRunnerUseCase interface {
		Execute(ctx context.Context, input CreateRunnerInput) (CreateRunnerOutput, error)
	}

	// CreateRunnerInput input data
	CreateRunnerInput struct {
		UserID         uuid.UUID `json:"-"`
		MonitorID      uuid.UUID `json:"monitor_id" binding:"required"`
		Name           string    `json:"name" binding:"required"`
		Region         string    `json:"region" binding:"required"`
		IntervalSecond int       `json:"interval_second" binding:"required"`
	}

	// CreateRunnerPresenter output port
	CreateRunnerPresenter interface {
		Output(*domain.Runner) CreateRunnerOutput
	}

	// CreateRunnerInput output data
	CreateRunnerOutput struct {
		ID             uuid.UUID `json:"id"`
		UserID         uuid.UUID `json:"user_id"`
		MonitorID      uuid.UUID `json:"monitor_id"`
		Name           string    `json:"name"`
		Region         string    `json:"region"`
		IntervalSecond int       `json:"interval_second"`
		IsEnabled      bool      `json:"is_enabled"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	createRunnerInteractor struct {
		repo      domain.RunnerRepository
		presenter CreateRunnerPresenter
	}
)

func NewCreateRunnerInteractor(
	repo domain.RunnerRepository,
	presenter CreateRunnerPresenter,
) CreateRunnerUseCase {
	return &createRunnerInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *createRunnerInteractor) Execute(ctx context.Context, input CreateRunnerInput) (CreateRunnerOutput, error) {
	runner, err := domain.NewRunner(
		input.UserID,
		input.MonitorID,
		input.Name,
		input.Region,
		input.IntervalSecond,
	)
	if err != nil {
		return CreateRunnerOutput{}, err
	}

	created, err := i.repo.Create(ctx, *runner)
	if err != nil {
		return CreateRunnerOutput{}, err
	}

	return i.presenter.Output(created), nil
}
