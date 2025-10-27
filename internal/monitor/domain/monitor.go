package domain

import (
	"context"
	"easy-go-monitor/internal/codes"
	"net/url"
	"time"

	"github.com/google/uuid"
)

type (
	MonitorRepository interface {
		Create(ctx context.Context, monitor Monitor) (*Monitor, error)
		FindByID(ctx context.Context, id uuid.UUID) (*Monitor, error)
		FindAll(ctx context.Context, userID uuid.UUID) ([]*Monitor, error)
		Update(ctx context.Context, monitor Monitor) (*Monitor, error)
		Delete(ctx context.Context, id uuid.UUID) error
	}

	Monitor struct {
		ID        uuid.UUID
		UserID    uuid.UUID
		Name      string
		URL       string
		Type      string
		Settings  map[string]any
		IsEnabled bool
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	MonitorSettings struct {
		Method    string            `json:"method"`
		TimeoutMs int               `json:"timeout_ms"`
		Headers   map[string]string `json:"headers"`
		Body      string            `json:"body"`
	}
)

// NewMonitor creates a new Monitor entity with validation.
func NewMonitor(
	userID uuid.UUID,
	name string,
	monitorUrl string,
	monitorType string,
	monitorSettings map[string]any,
) (*Monitor, error) {
	if userID == uuid.Nil {
		return nil, codes.ErrInvalidUUID
	}
	if name == "" {
		return nil, codes.ErrInvalidMonitorName
	}
	if _, err := url.ParseRequestURI(monitorUrl); err != nil {
		return nil, codes.ErrInvalidMonitorURL
	}
	if monitorType != "http" && monitorType != "tcp" && monitorType != "ping" {
		return nil, codes.ErrInvalidMonitorType
	}
	if monitorSettings == nil {
		monitorSettings = make(map[string]any)
	}

	return &Monitor{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      name,
		URL:       monitorUrl,
		Type:      monitorType,
		Settings:  monitorSettings,
		IsEnabled: true,
	}, nil
}
