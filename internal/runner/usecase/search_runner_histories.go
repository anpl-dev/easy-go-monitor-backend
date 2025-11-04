package usecase

import (
	"context"
	"easy-go-monitor/internal/runner/domain"

	"github.com/google/uuid"
)

type (
	SearchRunnerHistoriesUseCase interface {
		Execute(ctx context.Context, input SearchRunnerHistoriesInput) ([]SearchRunnerHistoriesOutput, error)
	}

	SearchRunnerHistoriesInput struct {
		UserID  uuid.UUID `json:"-"`
		Status  string    `json:"status"`
		Minutes int32     `json:"minutes"`
	}

	SearchRunnerHistoriesPresenter interface {
		Output([]domain.RunnerHistory) []SearchRunnerHistoriesOutput
	}

	SearchRunnerHistoriesOutput struct {
		ID             uuid.UUID `json:"id"`
		Status         string    `json:"status"`
		Message        string    `json:"message"`
		StartedAt      string    `json:"started_at"`
		EndedAt        string    `json:"ended_at"`
		ResponseTimeMs int32     `json:"response_time_ms"`
	}

	findRunnerFailHistoriesInteractor struct {
		repo      domain.RunnerHistoryRepository
		presenter SearchRunnerHistoriesPresenter
	}
)

func NewSearchRunnerHistoriesInteractor(
	repo domain.RunnerHistoryRepository,
	presenter SearchRunnerHistoriesPresenter,
) SearchRunnerHistoriesUseCase {
	return &findRunnerFailHistoriesInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *findRunnerFailHistoriesInteractor) Execute(ctx context.Context, input SearchRunnerHistoriesInput) ([]SearchRunnerHistoriesOutput, error) {
	histories, err := i.repo.Search(ctx, input.UserID, input.Status, int(input.Minutes))
	if err != nil {
		return nil, err
	}

	values := make([]domain.RunnerHistory, len(histories))
	for idx, h := range histories {
		values[idx] = *h
	}

	return i.presenter.Output(values), nil
}
