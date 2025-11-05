package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	RunnerHistoryRepository interface {
		Save(ctx context.Context, history RunnerHistory) error
		FindByID(ctx context.Context, runnerID uuid.UUID) ([]*RunnerHistory, error)
		Search(ctx context.Context, userID uuid.UUID, status string, minutes int) ([]*RunnerHistory, error)
	}

	RunnerHistory struct {
		ID             uuid.UUID
		RunnerID       uuid.UUID
		RunnerName     string
		Status         string
		Message        *string
		StartedAt      time.Time
		EndedAt        *time.Time
		ResponseTimeMs *int32
		CreatedAt      time.Time
	}
)
