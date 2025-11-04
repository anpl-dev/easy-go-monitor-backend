package presenter

import (
	"easy-go-monitor/internal/runner/domain"
	"easy-go-monitor/internal/runner/usecase"
)

type UpdateRunnerPresenter struct{}

func NewUpdateRunnerPresenter() *UpdateRunnerPresenter {
	return &UpdateRunnerPresenter{}
}

func (p *UpdateRunnerPresenter) Output(runner *domain.Runner) usecase.UpdateRunnerOutput {
	if runner == nil {
		return usecase.UpdateRunnerOutput{}
	}
	return usecase.UpdateRunnerOutput{
		ID:             runner.ID,
		UserID:         runner.UserID,
		MonitorID:      runner.MonitorID,
		Name:           runner.Name,
		Region:         runner.Region,
		IntervalSecond: runner.IntervalSecond,
		IsEnabled:      runner.IsEnabled,
		UpdatedAt:      runner.UpdatedAt,
	}
}
