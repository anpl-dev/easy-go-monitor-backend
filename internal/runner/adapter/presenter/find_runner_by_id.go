package presenter

import (
	"easy-go-monitor/internal/runner/domain"
	"easy-go-monitor/internal/runner/usecase"
)

type FindRunnerByIDPresenter struct{}

func NewFindRunnerByIDPresenter() *FindRunnerByIDPresenter {
	return &FindRunnerByIDPresenter{}
}

func (p *FindRunnerByIDPresenter) Output(runner *domain.Runner) usecase.FindRunnerByIDOutput {
	if runner == nil {
		return usecase.FindRunnerByIDOutput{}
	}
	return usecase.FindRunnerByIDOutput{
		ID:             runner.ID,
		UserID:         runner.UserID,
		MonitorID:      runner.MonitorID,
		Name:           runner.Name,
		Region:         runner.Region,
		IntervalSecond: runner.IntervalSecond,
		IsEnabled:      runner.IsEnabled,
		CreatedAt:      runner.CreatedAt,
		UpdatedAt:      runner.UpdatedAt,
	}
}
