package repository

import (
	"context"
	"database/sql"
	"easy-go-monitor/db/sqlcgen"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/runner/domain"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RunnerPostgresRepository struct {
	queries *sqlcgen.Queries
}

func NewRunnerPostgresRepository(pool *pgxpool.Pool) *RunnerPostgresRepository {
	return &RunnerPostgresRepository{queries: sqlcgen.New(pool)}
}

func toDomainRunner(s sqlcgen.Runner) *domain.Runner {
	return &domain.Runner{
		ID:             s.ID,
		UserID:         s.UserID,
		MonitorID:      s.MonitorID,
		Name:           s.Name,
		Region:         s.Region,
		IntervalSecond: int(s.IntervalSecond),
		IsActive:       *s.IsActive,
		CreatedAt:      s.CreatedAt.Time,
		UpdatedAt:      s.UpdatedAt.Time,
	}
}

func (r *RunnerPostgresRepository) Create(ctx context.Context, runner domain.Runner) (*domain.Runner, error) {
	row, err := r.queries.CreateRunner(ctx, sqlcgen.CreateRunnerParams{
		ID:             runner.ID,
		UserID:         runner.UserID,
		MonitorID:      runner.MonitorID,
		Name:           runner.Name,
		Region:         runner.Region,
		IntervalSecond: int32(runner.IntervalSecond),
		IsActive:       &runner.IsActive,
	})
	if err != nil {
		return nil, err
	}
	return toDomainRunner(row), nil
}

func (r *RunnerPostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Runner, error) {
	row, err := r.queries.FindRunnerByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, codes.ErrNotFound
		}
		return nil, err
	}
	return toDomainRunner(row), nil
}
