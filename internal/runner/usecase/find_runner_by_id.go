package usecase

import (
	"context"
	"easy-go-monitor/internal/runner/domain"
	"time"

	"github.com/google/uuid"
)

type (
	// FindRunnerByIDUseCase input port
	FindRunnerByIDUseCase interface {
		Execute(ctx context.Context, input FindRunnerByIDInput) (FindRunnerByIDOutput, error)
	}

	// FindRunnerByIDInput input data
	FindRunnerByIDInput struct {
		ID uuid.UUID `json:"-"`
	}

	// FindRunnerByIDPresenter output port
	FindRunnerByIDPresenter interface {
		Output(*domain.Runner) FindRunnerByIDOutput
	}

	// FindRunnerByIDOutput output data
	FindRunnerByIDOutput struct {
		ID             uuid.UUID `json:"id"`
		UserID         uuid.UUID `json:"user_id"`
		MonitorID      uuid.UUID `json:"runner_id"`
		Name           string    `json:"name"`
		Region         string    `json:"region"`
		IntervalSecond int       `json:"interval_second"`
		IsEnabled      bool      `json:"is_enabled"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	findRunnerByIDInteractor struct {
		repo      domain.RunnerRepository
		presenter FindRunnerByIDPresenter
	}
)

func NewFindRunnerByIDInteractor(
	repo domain.RunnerRepository,
	presenter FindRunnerByIDPresenter,
) FindRunnerByIDUseCase {
	return &findRunnerByIDInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *findRunnerByIDInteractor) Execute(ctx context.Context, input FindRunnerByIDInput) (FindRunnerByIDOutput, error) {
	runner, err := i.repo.FindByID(ctx, input.ID)
	if err != nil {
		return FindRunnerByIDOutput{}, err
	}
	if runner == nil {
		return FindRunnerByIDOutput{}, nil
	}
	return i.presenter.Output(runner), nil
}
