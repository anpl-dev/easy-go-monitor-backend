package usecase

import (
	"context"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/runner/domain"
	"time"

	"github.com/google/uuid"
)

type (
	// FindAllRunnersUseCase input port
	FindAllRunnersUseCase interface {
		Execute(ctx context.Context, input FindAllRunnersInput) ([]FindAllRunnersOutput, error)
	}

	// FindAllRunnersInput input data
	FindAllRunnersInput struct {
		UserID uuid.UUID `json:"-"`
	}

	// FindAllRunnersPresenter output port
	FindAllRunnersPresenter interface {
		Output([]*domain.Runner) []FindAllRunnersOutput
	}

	// FindAllRunnersOutput output data
	FindAllRunnersOutput struct {
		ID             uuid.UUID `json:"id"`
		UserID         uuid.UUID `json:"user_id"`
		MonitorID      uuid.UUID `json:"monitor_id"`
		Name           string    `json:"name"`
		Region         string    `json:"region"`
		IntervalSecond int       `json:"interval_second"`
		IsEnabled      bool      `json:"is_enabled"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	findAllRunnersInteractor struct {
		repo      domain.RunnerRepository
		presenter FindAllRunnersPresenter
	}
)

func NewFindAllRunnersInteractor(
	repo domain.RunnerRepository,
	presenter FindAllRunnersPresenter,
) FindAllRunnersUseCase {
	return &findAllRunnersInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (i *findAllRunnersInteractor) Execute(ctx context.Context, input FindAllRunnersInput) ([]FindAllRunnersOutput, error) {
	if input.UserID != uuid.Nil {
		return nil, codes.ErrInvalidUUID
	}

	runners, err := i.repo.FindAll(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	return i.presenter.Output(runners), nil
}
