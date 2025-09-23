package repository

import (
	"context"
	"database/sql"
	"go-monitor-tool/internal/adapter/repository/sqlcgen"
	"go-monitor-tool/internal/domain"
)

type MonitorPostgresRepository struct {
	q *sqlcgen.Queries
}

func NewMonitorPostgresRepository(db *sql.DB) *MonitorPostgresRepository {
	return &MonitorPostgresRepository{q: sqlcgen.New(db)}
}

func (r *MonitorPostgresRepository) Create(ctx context.Context, m domain.Monitor) (*domain.Monitor, error) {
	row, err := r.q.CreateMonitor(ctx, sqlcgen.CreateMonitorParams{
		ID: m.ID,
		UserID: m.UserID,
		Name: m.Name,
		Url: m.Url, 
		IntervalSecond: int32(m.IntervalSecond),
	})
	if err != nil {
		return nil, err
	}
	return &domain.Monitor{
		ID: row.ID,
		UserID: row.UserID,
		Name: row.Name,
		Url: row.Url,
		IntervalSecond: int(row.IntervalSecond),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}, nil
}