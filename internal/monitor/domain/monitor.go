package domain

import (
	"context"
	"easy-go-monitor/internal/codes"
	"net/url"
	"time"

	"github.com/google/uuid"
)

const (
	MonitorTypeHTTP = "http"
	MonitorTypeTCP  = "tcp"
	MonitorTypePing = "ping"
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
		Settings  MonitorSettings
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
	mUrl string,
	mType string,
	settings MonitorSettings,
) (*Monitor, error) {
	if userID == uuid.Nil {
		return nil, codes.ErrInvalidUUID
	}
	if name == "" {
		return nil, codes.ErrInvalidMonitorName
	}
	if _, err := url.ParseRequestURI(mUrl); err != nil {
		return nil, codes.ErrInvalidMonitorURL
	}
	if mType != "http" && mType != "tcp" && mType != "ping" {
		return nil, codes.ErrInvalidMonitorType
	}
	switch mType {
	case MonitorTypeHTTP, MonitorTypeTCP, MonitorTypePing:
		// OK
	default:
		return nil, codes.ErrInvalidMonitorType
	}

	return &Monitor{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      name,
		URL:       mUrl,
		Type:      mType,
		Settings:  settings,
		IsEnabled: true,
	}, nil
}

func NewMonitorSettings(
	method string,
	timeoutMS int,
	headers map[string]string,
	body string,
) (*MonitorSettings, error) {
	if method == "" {
		return nil, codes.ErrInvalidMonitorMethod
	}
	if timeoutMS <= 0 {
		timeoutMS = 5000
	}
	return &MonitorSettings{
		Method:    method,
		TimeoutMs: timeoutMS,
		Headers:   headers,
		Body:      body,
	}, nil
}
