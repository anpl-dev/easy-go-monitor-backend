package usecase

import (
	"context"
	"easy-go-monitor/internal/runner/domain"

	"github.com/google/uuid"
)

type (
	FindRunnerHistoriesUseCase interface {
		Execute(ctx context.Context, input FindRunnerHistoriesInput) ([]FindRunnerHistoriesOutput, error)
	}

	FindRunnerHistoriesInput struct {
		RunnerID uuid.UUID
	}

	FindRunnerHistoriesPresenter interface {
		Output([]domain.RunnerHistory) []FindRunnerHistoriesOutput
	}

	FindRunnerHistoriesOutput struct {
		ID             uuid.UUID `json:"id"`
		RunnerID       uuid.UUID `json:"runner_id"`
		RunnerName     string    `json:"runner_name"`
		Status         string    `json:"status"`
		Message        string    `json:"message"`
		StartedAt      string    `json:"started_at"`
		EndedAt        string    `json:"ended_at"`
		ResponseTimeMs int32     `json:"response_time_ms"`
	}

	findRunnerHistoriesInteractor struct {
		repo      domain.RunnerHistoryRepository
		presenter FindRunnerHistoriesPresenter
	}
)

func NewFindRunnerHistoriesInteractor(
	repo domain.RunnerHistoryRepository,
	presenter FindRunnerHistoriesPresenter,
) FindRunnerHistoriesUseCase {
	return &findRunnerHistoriesInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *findRunnerHistoriesInteractor) Execute(ctx context.Context, input FindRunnerHistoriesInput) ([]FindRunnerHistoriesOutput, error) {
	histories, err := i.repo.FindByID(ctx, input.RunnerID)
	if err != nil {
		return nil, err
	}

	values := make([]domain.RunnerHistory, len(histories))
	for idx, h := range histories {
		values[idx] = *h
	}

	return i.presenter.Output(values), nil
}
