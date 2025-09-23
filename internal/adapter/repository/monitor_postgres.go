package repository

import (
	"context"
	"database/sql"
	"go-monitor-tool/internal/adapter/repository/sqlcgen"
	"go-monitor-tool/internal/domain"

	"github.com/google/uuid"
)

type MonitorPostgresRepository struct {
	queries *sqlcgen.Queries
}

func NewMonitorPostgresRepository(db *sql.DB) *MonitorPostgresRepository {
	return &MonitorPostgresRepository{queries: sqlcgen.New(db)}
}
func toDomainMonitor(s sqlcgen.Monitor) *domain.Monitor {
	return &domain.Monitor{
		ID:             s.ID,
		UserID:         s.UserID,
		Name:           s.Name,
		URL:            s.Url,
		IntervalSecond: int(s.IntervalSecond),
	}
}

func toSQLCMonitorParams(m domain.Monitor) sqlcgen.CreateMonitorParams {
	return sqlcgen.CreateMonitorParams{
		ID:             m.ID,
		UserID:         m.UserID,
		Name:           m.Name,
		Url:            m.URL,
		IntervalSecond: int32(m.IntervalSecond),
	}
}

func (r *MonitorPostgresRepository) Create(ctx context.Context, m domain.Monitor) (*domain.Monitor, error) {
	row, err := r.queries.CreateMonitor(ctx, toSQLCMonitorParams(m))
	if err != nil {
		return nil, err
	}
	return toDomainMonitor(row), nil
}

func (r *MonitorPostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Monitor, error) {
	row, err := r.queries.FindMonitorByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toDomainMonitor(row), nil
}

func (r *MonitorPostgresRepository) FindByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Monitor, error) {
	rows, err := r.queries.FindMonitorsByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Monitor, 0, len(rows))
	for _, row := range rows {
		result = append(result, toDomainMonitor(row))
	}
	return result, nil
}
