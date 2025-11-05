package usecase

import (
	"context"
	"easy-go-monitor/internal/runner/domain"
	"time"

	"github.com/google/uuid"
)

type (
	// UpdateRunnerUseCase input port
	UpdateRunnerUseCase interface {
		Execute(ctx context.Context, input UpdateRunnerInput) (UpdateRunnerOutput, error)
	}

	// UpdateRunnerInput input data
	UpdateRunnerInput struct {
		ID             uuid.UUID `json:"-"`
		MonitorID      uuid.UUID `json:"monitor_id"`
		Name           string    `json:"name"`
		Region         string    `json:"region"`
		IntervalSecond int       `json:"interval_second"`
		IsEnabled      bool      `json:"is_enabled"`
	}

	// UpdateRunnerPresenter output port
	UpdateRunnerPresenter interface {
		Output(*domain.Runner) UpdateRunnerOutput
	}

	// UpdateRunnerInput output data
	UpdateRunnerOutput struct {
		ID             uuid.UUID `json:"id"`
		UserID         uuid.UUID `json:"user_id"`
		MonitorID      uuid.UUID `json:"monitor_id"`
		Name           string    `json:"name"`
		Region         string    `json:"region"`
		IntervalSecond int       `json:"interval_second"`
		IsEnabled      bool      `json:"is_enabled"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	updateRunnerInteractor struct {
		repo      domain.RunnerRepository
		presenter UpdateRunnerPresenter
	}
)

func NewUpdateRunnerInteractor(
	repo domain.RunnerRepository,
	presenter UpdateRunnerPresenter,
) UpdateRunnerUseCase {
	return &updateRunnerInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *updateRunnerInteractor) Execute(ctx context.Context, input UpdateRunnerInput) (UpdateRunnerOutput, error) {
	runner := domain.Runner{
		ID:             input.ID,
		MonitorID:      input.MonitorID,
		Name:           input.Name,
		Region:         input.Region,
		IntervalSecond: input.IntervalSecond,
		IsEnabled:      input.IsEnabled,
	}
	updated, err := i.repo.Update(ctx, runner)
	if err != nil {
		return UpdateRunnerOutput{}, err
	}

	return i.presenter.Output(updated), nil
}
