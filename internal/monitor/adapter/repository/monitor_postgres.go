package repository

import (
	"context"
	"easy-go-monitor/db/sqlcgen"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/monitor/domain"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MonitorPostgresRepository struct {
	queries *sqlcgen.Queries
}

func NewMonitorPostgresRepository(pool *pgxpool.Pool) *MonitorPostgresRepository {
	return &MonitorPostgresRepository{queries: sqlcgen.New(pool)}
}
func toDomainMonitor(s sqlcgen.Monitor) *domain.Monitor {
	var settings domain.MonitorSettings
	if s.Settings != nil {
		if err := json.Unmarshal(s.Settings, &settings); err != nil {
			// Guard settings
			settings = domain.MonitorSettings{}
		}
	}
	// Null Safe
	isEnabled := false
	if s.IsEnabled {
		isEnabled = s.IsEnabled
	}
	return &domain.Monitor{
		ID:        s.ID,
		UserID:    s.UserID,
		Name:      s.Name,
		URL:       s.Url,
		Type:      s.Type,
		Settings:  &settings,
		IsEnabled: isEnabled,
		CreatedAt: s.CreatedAt.Time,
		UpdatedAt: s.UpdatedAt.Time,
	}
}

// Create
func (r *MonitorPostgresRepository) Create(ctx context.Context, m domain.Monitor) (*domain.Monitor, error) {
	settingsJSON, err := json.Marshal(m.Settings)
	if err != nil {
		return nil, codes.ErrJSONRequest
	}
	isEnalbled := true
	row, err := r.queries.CreateMonitor(ctx, sqlcgen.CreateMonitorParams{
		ID:        m.ID,
		UserID:    m.UserID,
		Name:      m.Name,
		Url:       m.URL,
		Type:      m.Type,
		IsEnabled: isEnalbled,
		Settings:  settingsJSON,
	})
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == codes.PostgresForeignKeyViolation {
				return nil, codes.ErrNotFound
			}
			if pgErr.Code == codes.PostgresUniqueViolation {
				return nil, codes.ErrAlreadyExists
			}
		}
		return nil, err
	}
	return toDomainMonitor(row), nil
}

// FindByID
func (r *MonitorPostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Monitor, error) {
	row, err := r.queries.FindMonitorByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, codes.ErrNotFound
		}
		return nil, err
	}
	return toDomainMonitor(row), nil
}

// FindAll
func (r *MonitorPostgresRepository) FindAll(ctx context.Context, userID uuid.UUID) ([]*domain.Monitor, error) {
	rows, err := r.queries.FindAllMonitors(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, codes.ErrNotFound
		}
		return []*domain.Monitor{}, nil
	}

	result := make([]*domain.Monitor, 0, len(rows))
	for _, row := range rows {
		result = append(result, toDomainMonitor(row))
	}
	return result, nil
}

// Update
func (r *MonitorPostgresRepository) Update(ctx context.Context, m domain.Monitor) (*domain.Monitor, error) {
	settingsJSON, err := json.Marshal(m.Settings)
	if err != nil {
		return nil, codes.ErrJSONRequest
	}

	row, err := r.queries.UpdateMonitor(ctx, sqlcgen.UpdateMonitorParams{
		ID:        m.ID,
		Name:      m.Name,
		Url:       m.URL,
		Settings:  settingsJSON,
		IsEnabled: m.IsEnabled,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, codes.ErrNotFound
		}
		return nil, err
	}
	return toDomainMonitor(row), nil
}

// Delete
func (r *MonitorPostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.queries.DeleteMonitor(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return codes.ErrNotFound
		}
		return err
	}
	return nil
}
