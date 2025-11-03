package repository

import (
	"context"
	"easy-go-monitor/db/sqlcgen"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/runner/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RunnerHistoryPostgresRepository struct {
	queries *sqlcgen.Queries
}

func NewRunnerHistoryPostgresRepository(pool *pgxpool.Pool) *RunnerHistoryPostgresRepository {
	return &RunnerHistoryPostgresRepository{queries: sqlcgen.New(pool)}
}

func toDomainRunnerHistory(s sqlcgen.RunnerHistory) *domain.RunnerHistory {
	h := &domain.RunnerHistory{
		ID:             s.ID,
		RunnerID:       s.RunnerID,
		Status:         s.Status,
		Message:        s.Message,
		StartedAt:      *s.StartedAt,
		EndedAt:        s.EndedAt,
		ResponseTimeMs: s.ResponseTimeMs,
		CreatedAt:      *s.CreatedAt,
	}
	if s.StartedAt != nil {
		h.StartedAt = *s.StartedAt
	}
	if s.CreatedAt != nil {
		h.CreatedAt = *s.CreatedAt
	}
	return h
}

func (r *RunnerHistoryPostgresRepository) Save(ctx context.Context, h domain.RunnerHistory) error {
	params := sqlcgen.SaveRunnerHistoryParams{
		ID:             h.ID,
		RunnerID:       h.RunnerID,
		Status:         h.Status,
		Message:        h.Message,
		StartedAt:      &h.StartedAt,
		EndedAt:        h.EndedAt,
		ResponseTimeMs: h.ResponseTimeMs,
	}
	if err := r.queries.SaveRunnerHistory(ctx, params); err != nil {
		return codes.Wrap(codes.ErrInternal, err)
	}
	return nil
}

func (r *RunnerHistoryPostgresRepository) FindHistory(ctx context.Context, runnerID uuid.UUID) ([]*domain.RunnerHistory, error) {
	rows, err := r.queries.FindHistoryByRunnerID(ctx, runnerID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return []*domain.RunnerHistory{}, nil
		}
		return nil, codes.Wrap(codes.ErrInternal, err)
	}

	result := make([]*domain.RunnerHistory, 0, len(rows))
	for _, row := range rows {
		result = append(result, toDomainRunnerHistory(row))
	}
	return result, nil
}
