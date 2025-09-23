package domain

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidMonitorName     = errors.New("invalid monitor name")
	ErrInvalidMonitorURL      = errors.New("invalid monitor url")
	ErrInvalidMonitorInterval = errors.New("monitor interval must be greater than zero")
)

type (
	MonitorRepository interface {
		Create(ctx context.Context, m Monitor) (*Monitor, error)
		FindByID(ctx context.Context, id uuid.UUID) (*Monitor, error)
		FindByUserID(ctx context.Context, user_id uuid.UUID) ([]*Monitor, error)
		// Update(ctx context.Context, m Monitor) (*Monitor, error)
		// Delete(ctx context.Context, id uuid.UUID) error
	}

	Monitor struct {
		ID             uuid.UUID
		UserID         uuid.UUID
		Name           string
		Url            string
		IntervalSecond int
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
)

// NewMonitor creates a new Monitor entity with validation.
func NewMonitor(userID uuid.UUID, name string, rawURL string, interval int) (*Monitor, error) {
	if name == "" {
		return nil, ErrInvalidMonitorName
	}
	if interval <= 0 {
		return nil, ErrInvalidMonitorInterval
	}
	if _, err := url.ParseRequestURI(rawURL); err != nil {
		return nil, ErrInvalidMonitorURL
	}

	now := time.Now()
	return &Monitor{
		ID:             uuid.New(),
		UserID:         userID,
		Name:           name,
		Url:            rawURL,
		IntervalSecond: interval,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func (m *Monitor) Touch() {
	m.UpdatedAt = time.Now()
}
