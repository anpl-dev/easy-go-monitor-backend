package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	RunnerHistoryRepository interface {
		Save(ctx context.Context, history RunnerHistory) error
		FindByRunnerID(ctx context.Context, runnerID uuid.UUID) ([]*RunnerHistory, error)
	}

	RunnerHistory struct {
		ID             uuid.UUID  `json:"id"`
		RunnerID       uuid.UUID  `json:"runner_id"`
		Status         string     `json:"status"`
		Message        *string    `json:"message"`
		StartedAt      time.Time  `json:"started_at"`
		EndedAt        *time.Time `json:"ended_at"`
		DurationMs     *int32     `json:"duration_ms"`
		ResponseTimeMs *int32     `json:"response_time_ms"`
		CreatedAt      time.Time  `json:"created_at"`
	}
)
