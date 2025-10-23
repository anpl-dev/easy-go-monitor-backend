package domain

import (
	"context"
	"easy-go-monitor/internal/constraint"
	"time"

	"github.com/google/uuid"
)

type (
	RunnerRepository interface {
		Create(ctx context.Context, runner Runner) (*Runner, error)
		FindByID(ctx context.Context, id uuid.UUID) (*Runner, error)
		Update(ctx context.Context, r Runner) (*Runner, error)
		Delete(ctx context.Context, id uuid.UUID) error
	}

	Runner struct {
		ID             uuid.UUID
		UserID         uuid.UUID
		MonitorID      uuid.UUID
		Name           string
		Region         string
		IntervalSecond int
		IsActive       bool
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
)

func NewRunner(
	userID uuid.UUID,
	monitorID uuid.UUID,
	name string,
	interval_second int,
	region string,
) (*Runner, error) {
	if userID == uuid.Nil {
		return nil, constraint.ErrInvalidUUID
	}
	if monitorID == uuid.Nil {
		return nil, constraint.ErrInvalidUUID
	}
	if name == "" {
		return nil, constraint.ErrInvalidRunnerName
	}
	if region == "" {
		return nil, constraint.ErrInvalidRunnerRegion
	}
	if interval_second <= 0 {
		return nil, constraint.ErrInvalidRunnerInterval
	}

	return &Runner{
		ID:             uuid.New(),
		UserID:         userID,
		MonitorID:      monitorID,
		Name:           name,
		Region:         region,
		IntervalSecond: interval_second,
		IsActive:       true,
	}, nil
}
