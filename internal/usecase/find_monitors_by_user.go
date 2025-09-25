package usecase

import (
	"context"
	"go-monitor-tool/internal/domain"
	"time"

	"github.com/google/uuid"
)

type (
	FindMonitorsByUserUseCase interface {
		Execute(ctx context.Context, input FindMonitorsByUserInput) ([]FindMonitorsByUserOutput, error)
	}

	FindMonitorsByUserInput struct {
		UserID uuid.UUID `json:"user_id" binding:"required"`
	}

	FindMonitorsByUserPresenter interface {
		Output([]*domain.Monitor) []FindMonitorsByUserOutput
	}

	FindMonitorsByUserOutput struct {
		ID             uuid.UUID `json:"id"`
		UserID         uuid.UUID `json:"user_id"`
		Name           string    `json:"name"`
		URL            string    `json:"url"`
		IntervalSecond int       `json:"interval_second"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	findMonitorsByUserInteractor struct {
		repo      domain.MonitorRepository
		presenter FindMonitorsByUserPresenter
	}
)

func NewFindMonitorsByUserInteractor(
	repo domain.MonitorRepository,
	presenter FindMonitorsByUserPresenter,
) FindMonitorsByUserUseCase {
	return &findMonitorsByUserInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *findMonitorsByUserInteractor) Execute(ctx context.Context, input FindMonitorsByUserInput) ([]FindMonitorsByUserOutput, error) {
	monitors, err := i.repo.FindByUserID(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	return i.presenter.Output(monitors), nil
}
