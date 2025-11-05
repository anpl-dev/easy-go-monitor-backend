package usecase

import (
	"context"
	"easy-go-monitor/internal/infra/logger"
	"easy-go-monitor/internal/runner/domain"
	"sync"

	"github.com/google/uuid"
)

type (
	ExecuteRunnerUseCase interface {
		Execute(ctx context.Context, input ExecuteRunnerInput) ([]ExecuteRunnerOutput, error)
	}

	ExecuteRunnerInput struct {
		RunnerIDs []uuid.UUID
	}

	ExecuteRunnerPresenter interface {
		Output([]domain.MonitorResult) []ExecuteRunnerOutput
	}
	ExecuteRunnerOutput struct {
		MonitorID string `json:"monitor_id"`
		Status    string `json:"status"`
		LatencyMs int64  `json:"latency_ms"`
	}

	executeRunnerInteractor struct {
		service   domain.RunnerService
		presenter ExecuteRunnerPresenter
		log       *logger.Logger
	}
)

func NewExecuteRunnerInteractor(
	service domain.RunnerService,
	presenter ExecuteRunnerPresenter,
	log *logger.Logger,
) ExecuteRunnerUseCase {
	return &executeRunnerInteractor{
		service:   service,
		presenter: presenter,
		log:       log,
	}
}

func (i *executeRunnerInteractor) Execute(ctx context.Context, input ExecuteRunnerInput) ([]ExecuteRunnerOutput, error) {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	allResults := []domain.MonitorResult{}

	for _, runnerID := range input.RunnerIDs {
		wg.Add(1)
		go func(id uuid.UUID) {
			i.log.Debug("Runner execution started", "runner_id", id)
			defer wg.Done()

			results, err := i.service.Run(ctx, id)
			if err != nil {
				i.log.Error("Runner execution failed", "runner_id", id, "error", err)
				return
			}

			mu.Lock()
			allResults = append(allResults, results...)
			mu.Unlock()

			i.log.Info("Runner execution completed", "runner_id", id, "result_count", len(results))
		}(runnerID)
	}

	wg.Wait()
	return i.presenter.Output(allResults), nil
}
