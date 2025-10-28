package usecase

import (
	"context"
	"easy-go-monitor/internal/monitor/domain"
	"time"

	"github.com/google/uuid"
)

type (
	// FindAllMonitorsUseCase input port
	FindAllMonitorsUseCase interface {
		Execute(ctx context.Context, input FindAllMonitorsInput) ([]FindAllMonitorsOutput, error)
	}

	// FindAllMonitorsInput input data
	FindAllMonitorsInput struct {
		UserID string `json:"-"`
	}

	// FindAllMonitorsPresenter output port
	FindAllMonitorsPresenter interface {
		Output([]*domain.Monitor) []FindAllMonitorsOutput
	}

	// FindAllMonitorsOutput output data
	FindAllMonitorsOutput struct {
		ID        uuid.UUID               `json:"id"`
		UserID    uuid.UUID               `json:"user_id"`
		Name      string                  `json:"name"`
		URL       string                  `json:"url"`
		Type      string                  `json:"type"`
		Settings  *domain.MonitorSettings `json:"settings"`
		IsEnabled bool                    `json:"is_enabled"`
		CreatedAt time.Time               `json:"created_at"`
		UpdatedAt time.Time               `json:"updated_at"`
	}

	findAllMonitorsInteractor struct {
		repo      domain.MonitorRepository
		presenter FindAllMonitorsPresenter
	}
)

func NewFindAllMonitorsInteractor(
	repo domain.MonitorRepository,
	presenter FindAllMonitorsPresenter,
) FindAllMonitorsUseCase {
	return &findAllMonitorsInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *findAllMonitorsInteractor) Execute(ctx context.Context, input FindAllMonitorsInput) ([]FindAllMonitorsOutput, error) {
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
