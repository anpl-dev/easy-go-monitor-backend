package usecase

import (
	"context"
	"easy-go-monitor/internal/runner/domain"

	"github.com/google/uuid"
)

type (
	FindRunnerHistoryUseCase interface {
		Execute(ctx context.Context, input FindRunnerHistoryInput) ([]FindRunnerHistoryOutput, error)
	}

	FindRunnerHistoryInput struct {
		RunnerID uuid.UUID
	}

	FindRunnerHistoryPresenter interface {
		Output([]domain.RunnerHistory) []FindRunnerHistoryOutput
	}

	FindRunnerHistoryOutput struct {
		ID             uuid.UUID `json:"id"`
		Status         string    `json:"status"`
		Message        string    `json:"message"`
		StartedAt      string    `json:"started_at"`
		EndedAt        string    `json:"ended_at"`
		ResponseTimeMs int32     `json:"response_time_ms"`
	}

	findRunnerHistoryInteractor struct {
		repo      domain.RunnerHistoryRepository
		presenter FindRunnerHistoryPresenter
	}
)

func NewFindRunnerHistoryInteractor(
	repo domain.RunnerHistoryRepository,
	presenter FindRunnerHistoryPresenter,
) FindRunnerHistoryUseCase {
	return &findRunnerHistoryInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *findRunnerHistoryInteractor) Execute(ctx context.Context, input FindRunnerHistoryInput) ([]FindRunnerHistoryOutput, error) {
	histories, err := i.repo.FindHistory(ctx, input.RunnerID)
	if err != nil {
		return nil, err
	}

	// convert []*domain.RunnerHistory → []domain.RunnerHistory
	values := make([]domain.RunnerHistory, len(histories))
	for idx, h := range histories {
		values[idx] = *h
	}

	return i.presenter.Output(values), nil
}
