package repository

import (
	"context"
	"easy-go-monitor/db/sqlcgen"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/infra/logger"
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
	logger  *logger.Logger
}

func NewMonitorPostgresRepository(pool *pgxpool.Pool, log *logger.Logger) *MonitorPostgresRepository {
	return &MonitorPostgresRepository{queries: sqlcgen.New(pool), logger: log}
}
func toDomainMonitor(s sqlcgen.Monitor) *domain.Monitor {
	var settings domain.MonitorSettings
	if s.Settings != nil {
		if err := json.Unmarshal(s.Settings, &settings); err != nil {
			settings = domain.MonitorSettings{}
		}
	}
	return &domain.Monitor{
		ID:        s.ID,
		UserID:    s.UserID,
		Name:      s.Name,
		URL:       s.Url,
		Type:      s.Type,
		Settings:  &settings,
		IsEnabled: s.IsEnabled,
		CreatedAt: *s.CreatedAt,
		UpdatedAt: *s.UpdatedAt,
	}
}

// Create
func (r *MonitorPostgresRepository) Create(ctx context.Context, m domain.Monitor) (*domain.Monitor, error) {
	r.logger.Info("Inserting monitor",
		"user_id", m.UserID.String(),
		"name", m.Name,
		"url", m.URL,
	)
	settingsJSON, err := json.Marshal(m.Settings)
	if err != nil {
		return nil, codes.ErrJSONRequest
	}
	row, err := r.queries.CreateMonitor(ctx, sqlcgen.CreateMonitorParams{
		ID:        m.ID,
		UserID:    m.UserID,
		Name:      m.Name,
		Url:       m.URL,
		Type:      m.Type,
		Settings:  settingsJSON,
		IsEnabled: m.IsEnabled,
	})
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == codes.PostgresForeignKeyViolation {
				return nil, codes.ErrInvalidMonitorRequest
			}
			if pgErr.Code == codes.PostgresUniqueViolation {
				return nil, codes.ErrAlreadyExists
			}
		}
		r.logger.Error("DB insert failed", "error", err)
		return nil, err
	}
	r.logger.Debug("Monitor inserted", "id", row.ID.String())
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
			return []*domain.Monitor{}, nil
		}
		return nil, err
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
		Type:      m.Type,
		Url:       m.URL,
		Settings:  settingsJSON,
		IsEnabled: m.IsEnabled,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, codes.ErrConflict
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

func (r *MonitorPostgresRepository) SetEnabled(
	ctx context.Context,
	id uuid.UUID,
	enabled bool,
) (*domain.Monitor, error) {
	row, err := r.queries.UpdateMonitorIsEnabled(ctx, sqlcgen.UpdateMonitorIsEnabledParams{
		ID:        id,
		IsEnabled: enabled,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, codes.ErrConflict
		}
		return nil, err
	}
	return toDomainMonitor(row), nil
}
