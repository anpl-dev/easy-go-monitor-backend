package presenter

import (
	"easy-go-monitor/internal/runner/domain"
	"easy-go-monitor/internal/runner/usecase"
)

type findRunnerHistoriesPresenter struct{}

func NewFindRunnerHistoriesPresenter() usecase.FindRunnerHistoriesPresenter {
	return &findRunnerHistoriesPresenter{}
}

func (p *findRunnerHistoriesPresenter) Output(histories []domain.RunnerHistory) []usecase.FindRunnerHistoriesOutput {
	result := make([]usecase.FindRunnerHistoriesOutput, len(histories))
	for i, h := range histories {
		result[i] = usecase.FindRunnerHistoriesOutput{
			ID:             h.ID,
			RunnerID:       h.RunnerID,
			RunnerName:     h.RunnerName,
			Status:         h.Status,
			Message:        *h.Message,
			StartedAt:      h.StartedAt.Format("2006-01-02 15:04:05"),
			EndedAt:        h.EndedAt.Format("2006-01-02 15:04:05"),
			ResponseTimeMs: *h.ResponseTimeMs,
		}
	}
	return result
}
