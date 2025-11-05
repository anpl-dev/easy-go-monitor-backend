package repository

import (
	"context"
	"easy-go-monitor/db/sqlcgen"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/infra/logger"
	"easy-go-monitor/internal/runner/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RunnerHistoryPostgresRepository struct {
	queries *sqlcgen.Queries
	log     *logger.Logger
}

func NewRunnerHistoryPostgresRepository(pool *pgxpool.Pool, log *logger.Logger) *RunnerHistoryPostgresRepository {
	return &RunnerHistoryPostgresRepository{
		queries: sqlcgen.New(pool),
		log:     log,
	}
}

func toDomainRunnerHistory(s sqlcgen.RunnerHistory) *domain.RunnerHistory {
	h := &domain.RunnerHistory{
		ID:             s.ID,
		RunnerID:       s.RunnerID,
		RunnerName:     s.RunnerName,
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
		RunnerName:     h.RunnerName,
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

func (r *RunnerHistoryPostgresRepository) FindByID(ctx context.Context, runnerID uuid.UUID) ([]*domain.RunnerHistory, error) {
	rows, err := r.queries.FindRunnerHistoriesByRunnerID(ctx, runnerID)
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

func (r *RunnerHistoryPostgresRepository) Search(
	ctx context.Context,
	userID uuid.UUID,
	status string,
	minutes int,
) ([]*domain.RunnerHistory, error) {
	rows, err := r.queries.SearchRunnerHistories(ctx, sqlcgen.SearchRunnerHistoriesParams{
		UserID:  userID,
		Status:  status,
		Minutes: int32(minutes),
	})
	if err == pgx.ErrNoRows {
		r.log.Debug("No Rows")
		return []*domain.RunnerHistory{}, codes.ErrNotFound
	}
	if err != nil {
		return nil, codes.Wrap(codes.ErrInternal, err)
	}
	result := make([]*domain.RunnerHistory, 0, len(rows))
	for _, row := range rows {
		result = append(result, toDomainRunnerHistory(row))
	}
	return result, nil
}
