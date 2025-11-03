package presenter

import (
	"easy-go-monitor/internal/runner/domain"
	"easy-go-monitor/internal/runner/usecase"
)

type FindAllRunnersPresenter struct{}

func NewFindAllRunnersPresenter() *FindAllRunnersPresenter {
	return &FindAllRunnersPresenter{}
}

func (p *FindAllRunnersPresenter) Output(runners []*domain.Runner) []usecase.FindAllRunnersOutput {
	if runners == nil {
		return []usecase.FindAllRunnersOutput{}
	}

	outputs := make([]usecase.FindAllRunnersOutput, 0, len(runners))
	for _, runner := range runners {
		if runner == nil {
			continue
		}
		outputs = append(outputs, usecase.FindAllRunnersOutput{
			ID:             runner.ID,
			UserID:         runner.UserID,
			MonitorID:      runner.MonitorID,
			Name:           runner.Name,
			Region:         runner.Region,
			IntervalSecond: runner.IntervalSecond,
			IsEnabled:      runner.IsEnabled,
			CreatedAt:      runner.CreatedAt,
			UpdatedAt:      runner.UpdatedAt,
		})
	}
	return outputs
}
