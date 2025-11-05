package domain

import (
	"context"
	"easy-go-monitor/internal/infra/logger"
	"easy-go-monitor/internal/monitor/domain"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type RunnerService interface {
	Run(ctx context.Context, runner_id uuid.UUID) ([]MonitorResult, error)
}

type runnerService struct {
	runnerRepo  RunnerRepository
	monitorRepo domain.MonitorRepository
	historyRepo RunnerHistoryRepository
	log         *logger.Logger
}

func NewRunnerService(
	runnerRepo RunnerRepository,
	monitorRepo domain.MonitorRepository,
	historyRepo RunnerHistoryRepository,
	log *logger.Logger,
) RunnerService {
	return &runnerService{
		runnerRepo:  runnerRepo,
		monitorRepo: monitorRepo,
		historyRepo: historyRepo,
		log:         log,
	}
}

func (s *runnerService) Run(ctx context.Context, runnerID uuid.UUID) ([]MonitorResult, error) {
	s.log.Debug("RunnerService: start running", "runner_id", runnerID)

	runner, err := s.runnerRepo.FindByID(ctx, runnerID)
	if err != nil {
		return []MonitorResult{}, err
	}

	monitor, err := s.monitorRepo.FindByID(ctx, runner.MonitorID)
	if err != nil {
		s.log.Warn("monitor not found for runner", "runner_id", runnerID)
		return []MonitorResult{}, nil
	}

	s.log.Debug("Running monitor", "monitor_id", monitor.ID, "url", monitor.URL)

	start := time.Now()
	status := "OK"
	message := ""

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, monitor.URL, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode >= 400 {
		status = "FAIL"
		if err != nil {
			message = err.Error()
		} else {
			message = res.Status
		}
	}

	if res != nil {
		res.Body.Close()
	}

	latency := time.Since(start)

	result := MonitorResult{
		MonitorID: monitor.ID,
		Status:    status,
		LatencyMs: latency.Milliseconds(),
		Timestamp: time.Now(),
	}

	endedAt := time.Now()
	responseTimeMs := int32(latency.Milliseconds())

	history := RunnerHistory{
		ID:             uuid.New(),
		RunnerID:       runnerID,
		RunnerName:     runner.Name,
		Status:         status,
		Message:        &message,
		StartedAt:      start,
		EndedAt:        &endedAt,
		ResponseTimeMs: &responseTimeMs,
		CreatedAt:      time.Now(),
	}

	if err := s.historyRepo.Save(ctx, history); err != nil {
		s.log.Error("failed to save runner history", "runner_id", runnerID, "error", err)
	}

	return []MonitorResult{result}, nil
}
