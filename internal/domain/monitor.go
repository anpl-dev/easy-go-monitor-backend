package domain

import (
	"context"
	"go-monitor-tool/internal/errors"
	"net/url"
	"time"

	"github.com/google/uuid"
)

type (
	MonitorRepository interface {
		Create(ctx context.Context, m Monitor) (*Monitor, error)
		FindByID(ctx context.Context, id uuid.UUID) (*Monitor, error)
		FindByUserID(ctx context.Context, user_id uuid.UUID) ([]*Monitor, error)
		Update(ctx context.Context, m Monitor) (*Monitor, error)
		Delete(ctx context.Context, id uuid.UUID) error
	}

	Monitor struct {
		ID             uuid.UUID
		UserID         uuid.UUID
		Name           string
		URL            string
		IntervalSecond int
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
)

// NewMonitor creates a new Monitor entity with validation.
func NewMonitor(userID uuid.UUID, name string, rawURL string, interval int) (*Monitor, error) {
	if name == "" {
		return nil, errors.ErrInvalidMonitorName
	}
	if interval <= 0 {
		return nil, errors.ErrInvalidMonitorInterval
	}
	if _, err := url.ParseRequestURI(rawURL); err != nil {
		return nil, errors.ErrInvalidMonitorURL
	}

	now := time.Now()
	return &Monitor{
		ID:             uuid.New(),
		UserID:         userID,
		Name:           name,
		URL:            rawURL,
		IntervalSecond: interval,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func (m *Monitor) Touch() {
	m.UpdatedAt = time.Now()
}
