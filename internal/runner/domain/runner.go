package domain

import (
	"context"
	"easy-go-monitor/internal/codes"
	"time"

	"github.com/google/uuid"
)

type (
	RunnerRepository interface {
		Create(ctx context.Context, runner Runner) (*Runner, error)
		FindByID(ctx context.Context, id uuid.UUID) (*Runner, error)
		FindAll(ctx context.Context, userID uuid.UUID) ([]*Runner, error)
		Update(ctx context.Context, runner Runner) (*Runner, error)
		Delete(ctx context.Context, id uuid.UUID) error
	}

	Runner struct {
		ID             uuid.UUID
		UserID         uuid.UUID
		MonitorID      uuid.UUID
		Name           string
		Region         string
		IntervalSecond int
		IsEnabled      bool
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
)

func NewRunner(
	userID uuid.UUID,
	monitorID uuid.UUID,
	name string,
	region string,
	interval_second int,
) (*Runner, error) {
	if userID == uuid.Nil {
		return nil, codes.ErrInvalidUUID
	}
	if monitorID == uuid.Nil {
		return nil, codes.ErrInvalidUUID
	}
	if name == "" {
		return nil, codes.ErrInvalidRunnerName
	}
	if region == "" {
		return nil, codes.ErrInvalidRunnerRegion
	}
	if interval_second <= 0 {
		return nil, codes.ErrInvalidRunnerInterval
	}

	return &Runner{
		ID:             uuid.New(),
		UserID:         userID,
		MonitorID:      monitorID,
		Name:           name,
		Region:         region,
		IntervalSecond: interval_second,
		IsEnabled:      true,
	}, nil
}
