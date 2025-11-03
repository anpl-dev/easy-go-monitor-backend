package presenter

import (
	"easy-go-monitor/internal/runner/domain"
	"easy-go-monitor/internal/runner/usecase"
)

type findRunnerHistoryPresenter struct{}

func NewFindRunnerHistoryPresenter() usecase.FindRunnerHistoryPresenter {
	return &findRunnerHistoryPresenter{}
}

func (p *findRunnerHistoryPresenter) Output(histories []domain.RunnerHistory) []usecase.FindRunnerHistoryOutput {
	result := make([]usecase.FindRunnerHistoryOutput, len(histories))
	for i, h := range histories {
		result[i] = usecase.FindRunnerHistoryOutput{
			ID:             h.ID,
			Status:         h.Status,
			Message:        *h.Message,
			StartedAt:      h.StartedAt.Format("2006-01-02 15:04:05"),
			EndedAt:        h.EndedAt.Format("2006-01-02 15:04:05"),
			DurationMs:     *h.DurationMs,
			ResponseTimeMs: *h.ResponseTimeMs,
		}
	}
	return result
}
