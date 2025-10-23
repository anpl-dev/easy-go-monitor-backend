package usecase

import (
	"context"
	"easy-go-monitor/internal/monitor/domain"
	"time"

	"github.com/google/uuid"
)

type (
	// SearchMonitorsUseCase input port
	SearchMonitorsUseCase interface {
		Execute(ctx context.Context, input SearchMonitorsInput) ([]SearchMonitorsOutput, error)
	}

	// SearchMonitorsInput input data
	SearchMonitorsInput struct {
		UserID string `json:"user_id,omitempty" binding:"uuid"`
		Name   string `json:"name,omitempty" `
	}

	// SearchMonitorsPresenter output port
	SearchMonitorsPresenter interface {
		Output([]*domain.Monitor) []SearchMonitorsOutput
	}

	// SearchMonitorsOutput output data
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
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, err
	}

	monitors, err := i.repo.FindAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	return i.presenter.Output(monitors), nil
}
