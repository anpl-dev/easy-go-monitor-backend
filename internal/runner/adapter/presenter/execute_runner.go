package presenter

import (
	"easy-go-monitor/internal/runner/domain"
	"easy-go-monitor/internal/runner/usecase"
)

type ExecuteRunnerPresenter struct{}

func NewExecuteRunnerPresenter() *ExecuteRunnerPresenter {
	return &ExecuteRunnerPresenter{}
}

func (p *ExecuteRunnerPresenter) Output(results []domain.MonitorResult) []usecase.ExecuteRunnerOutput {
	outputs := make([]usecase.ExecuteRunnerOutput, len(results))
	for i, r := range results {
		outputs[i] = usecase.ExecuteRunnerOutput{
			MonitorID: r.MonitorID.String(),
			Status:    r.Status,
			LatencyMs: r.LatencyMs,
		}
	}
	return outputs
}
