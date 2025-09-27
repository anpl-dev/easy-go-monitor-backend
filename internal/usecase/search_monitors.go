package usecase

import (
	"context"
	"go-monitor-tool/internal/domain"
	"time"

	"github.com/google/uuid"
)

type (
	SearchMonitorsUseCase interface {
		Execute(ctx context.Context, input SearchMonitorsInput) ([]SearchMonitorsOutput, error)
	}

	SearchMonitorsInput struct {
		UserID uuid.UUID `json:"user_id,omitempty" `
		Name   uuid.UUID `json:"name,omitempty" `
	}

	SearchMonitorsPresenter interface {
		Output([]*domain.Monitor) []SearchMonitorsOutput
	}

	SearchMonitorsOutput struct {
		ID             uuid.UUID `json:"id"`
		UserID         uuid.UUID `json:"user_id"`
		Name           string    `json:"name"`
		URL            string    `json:"url"`
		IntervalSecond int       `json:"interval_second"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	searchMonitorsInteractor struct {
		repo      domain.MonitorRepository
		presenter SearchMonitorsPresenter
	}
)

func NewSearchMonitorsInteractor(
	repo domain.MonitorRepository,
	presenter SearchMonitorsPresenter,
) SearchMonitorsUseCase {
	return &searchMonitorsInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *searchMonitorsInteractor) Execute(ctx context.Context, input SearchMonitorsInput) ([]SearchMonitorsOutput, error) {
	monitors, err := i.repo.FindByUserID(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	return i.presenter.Output(monitors), nil
}
