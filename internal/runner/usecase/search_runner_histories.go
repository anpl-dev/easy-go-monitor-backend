package usecase

import (
	"context"
	"easy-go-monitor/internal/infra/logger"
	"easy-go-monitor/internal/runner/domain"

	"github.com/google/uuid"
)

type (
	SearchRunnerHistoriesUseCase interface {
		Execute(ctx context.Context, input SearchRunnerHistoriesInput) ([]SearchRunnerHistoriesOutput, error)
	}

	SearchRunnerHistoriesInput struct {
		UserID  uuid.UUID `json:"-"`
		Status  string    `form:"status" binding:"required"`
		Minutes int32     `form:"minutes" binding:"required"`
	}

	SearchRunnerHistoriesPresenter interface {
		Output([]domain.RunnerHistory) []SearchRunnerHistoriesOutput
	}

	SearchRunnerHistoriesOutput struct {
		ID             uuid.UUID `json:"id"`
		RunnerID       uuid.UUID `json:"runner_id"`
		RunnerName     string    `json:"runner_name"`
		Status         string    `json:"status"`
		Message        string    `json:"message"`
		StartedAt      string    `json:"started_at"`
		EndedAt        string    `json:"ended_at"`
		ResponseTimeMs int32     `json:"response_time_ms"`
	}

	searchRunnerHistoriesInteractor struct {
		repo      domain.RunnerHistoryRepository
		presenter SearchRunnerHistoriesPresenter
		log       *logger.Logger
	}
)

func NewSearchRunnerHistoriesInteractor(
	repo domain.RunnerHistoryRepository,
	presenter SearchRunnerHistoriesPresenter,
	log *logger.Logger,
) SearchRunnerHistoriesUseCase {
	return &searchRunnerHistoriesInteractor{
		repo:      repo,
		presenter: presenter,
		log:       log,
	}
}

func (i *searchRunnerHistoriesInteractor) Execute(ctx context.Context, input SearchRunnerHistoriesInput) ([]SearchRunnerHistoriesOutput, error) {
	i.log.Info("SearchRunnerHistories Execute",
		"user_id", input.UserID.String(),
		"status", input.Status,
		"minutes", input.Minutes,
	)
	histories, err := i.repo.Search(ctx, input.UserID, input.Status, int(input.Minutes))

	if err != nil {
		i.log.Error("FindRunnerFailHistories repo error", "error", err)
		return nil, err
	}

	values := make([]domain.RunnerHistory, len(histories))
	for idx, h := range histories {
		values[idx] = *h
	}

	return i.presenter.Output(values), nil
}
