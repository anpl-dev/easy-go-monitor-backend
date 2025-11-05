package domain

import (
	"context"
	"easy-go-monitor/internal/codes"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	MonitorTypeHTTP = "HTTP"
)

type (
	MonitorRepository interface {
		Create(ctx context.Context, monitor Monitor) (*Monitor, error)
		FindByID(ctx context.Context, id uuid.UUID) (*Monitor, error)
		FindAll(ctx context.Context, userID uuid.UUID) ([]*Monitor, error)
		Update(ctx context.Context, monitor Monitor) (*Monitor, error)
		Delete(ctx context.Context, id uuid.UUID) error

		SetEnabled(ctx context.Context, id uuid.UUID, enabled bool) (*Monitor, error)
	}

	Monitor struct {
		ID        uuid.UUID
		UserID    uuid.UUID
		Name      string
		URL       string
		Type      string
		Settings  *MonitorSettings
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
) (*Monitor, error) {
	normalizedType := strings.ToUpper(monitorType)
	if userID == uuid.Nil {
		return nil, codes.ErrInvalidUUID
	}
	if name == "" {
		return nil, codes.ErrInvalidMonitorName
	}
	if monitorUrl == "" {
		return nil, codes.ErrInvalidMonitorURL
	}
	switch normalizedType {
	case MonitorTypeHTTP:
		// OK
	default:
		return nil, codes.ErrInvalidMonitorType
	}

	defaultSettings, err := NewMonitorSettingsByType(normalizedType)
	if err != nil {
		return nil, err
	}

	return &Monitor{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      name,
		URL:       monitorUrl,
		Type:      normalizedType,
		Settings:  defaultSettings,
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
	if headers == nil {
		headers = map[string]string{}
	}
	return &MonitorSettings{
		Method:    method,
		TimeoutMs: timeoutMS,
		Headers:   headers,
		Body:      body,
	}, nil
}

func NewMonitorSettingsByType(monitorType string) (*MonitorSettings, error) {
	switch monitorType {
	case MonitorTypeHTTP:
		return NewMonitorSettings("GET", 5000, nil, "")
	default:
		return nil, codes.ErrInvalidMonitorType
	}
}
