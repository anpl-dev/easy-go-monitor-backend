package presenter

import (
	"easy-go-monitor/internal/runner/domain"
	"easy-go-monitor/internal/runner/usecase"
)

type CreateRunnerPresenter struct{}

func NewCreateRunnerPresenter() *CreateRunnerPresenter {
	return &CreateRunnerPresenter{}
}

func (p *CreateRunnerPresenter) Output(runner *domain.Runner) usecase.CreateRunnerOutput {
	if runner == nil {
		return usecase.CreateRunnerOutput{}
	}
	return usecase.CreateRunnerOutput{
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
