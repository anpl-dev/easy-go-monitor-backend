package repository

import (
	"context"
	"database/sql"
	"easy-go-monitor/db/sqlcgen"
	"easy-go-monitor/internal/constraint"
	"easy-go-monitor/internal/monitor/domain"
	"errors"

	"github.com/google/uuid"
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
	return &domain.Monitor{
		ID:        s.ID,
		UserID:    s.UserID,
		Name:      s.Name,
		URL:       s.Url,
		CreatedAt: s.CreatedAt.Time,
		UpdatedAt: s.UpdatedAt.Time,
	}
}

func (r *MonitorPostgresRepository) Create(ctx context.Context, m domain.Monitor) (*domain.Monitor, error) {
	row, err := r.queries.CreateMonitor(ctx, sqlcgen.CreateMonitorParams{
		ID:     m.ID,
		UserID: m.UserID,
		Name:   m.Name,
		Url:    m.URL,
	})
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == constraint.PostgresForeignKeyViolation {
				return nil, constraint.ErrNotFound
			}
		}
		return nil, err
	}
	return toDomainMonitor(row), nil
}

func (r *MonitorPostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Monitor, error) {
	row, err := r.queries.FindMonitorByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, constraint.ErrNotFound
		}
		return nil, err
	}
	return toDomainMonitor(row), nil
}

func (r *MonitorPostgresRepository) FindByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Monitor, error) {
	rows, err := r.queries.FindAllMonitors(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, constraint.ErrNotFound
		}
		return nil, err
	}

	result := make([]*domain.Monitor, 0, len(rows))
	for _, row := range rows {
		result = append(result, toDomainMonitor(row))
	}
	return result, nil
}

func (r *MonitorPostgresRepository) Update(ctx context.Context, m domain.Monitor) (*domain.Monitor, error) {
	row, err := r.queries.UpdateMonitor(ctx, sqlcgen.UpdateMonitorParams{
		ID:   m.ID,
		Name: m.Name,
		Url:  m.URL,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, constraint.ErrNotFound
		}
		return nil, err
	}
	return toDomainMonitor(row), nil
}

func (r *MonitorPostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.queries.DeleteMonitor(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return constraint.ErrNotFound
		}
		return err
	}
	return nil
}
