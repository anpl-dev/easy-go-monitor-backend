package repository

import (
	"context"
	"easy-go-monitor/db/sqlcgen"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/runner/domain"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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
		IsEnabled:      s.IsEnabled,
		CreatedAt:      *s.CreatedAt,
		UpdatedAt:      *s.UpdatedAt,
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
		IsEnabled:      runner.IsEnabled,
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
	}
	return toDomainRunner(row), nil
}

func (r *RunnerPostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Runner, error) {
	row, err := r.queries.FindRunnerByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, codes.ErrNotFound
		}
		return nil, err
	}
	return toDomainRunner(row), nil
}

func (r *RunnerPostgresRepository) FindAll(ctx context.Context, userID uuid.UUID) ([]*domain.Runner, error) {
	rows, err := r.queries.FindAllRunners(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, codes.ErrNotFound
		}
		return nil, err
	}

	result := make([]*domain.Runner, 0, len(rows))
	for _, row := range rows {
		result = append(result, toDomainRunner(row))
	}
	return result, nil
}

func (r *RunnerPostgresRepository) Update(ctx context.Context, runner domain.Runner) (*domain.Runner, error) {
	row, err := r.queries.UpdateRunner(ctx, sqlcgen.UpdateRunnerParams{
		ID:             runner.ID,
		MonitorID:      runner.MonitorID,
		Name:           runner.Name,
		Region:         runner.Region,
		IntervalSecond: int32(runner.IntervalSecond),
		IsEnabled:      runner.IsEnabled,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, codes.ErrNotFound
		}
		return nil, err
	}
	return toDomainRunner(row), nil
}

func (r *RunnerPostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.queries.DeleteRunner(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return codes.ErrNotFound
		}
		return err
	}
	return nil
}
