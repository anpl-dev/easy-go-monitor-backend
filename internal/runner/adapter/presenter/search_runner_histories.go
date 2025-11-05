package presenter

import (
	"easy-go-monitor/internal/runner/domain"
	"easy-go-monitor/internal/runner/usecase"
)

type searchRunnerHistoriesPresenter struct{}

func NewSearchRunnerHistoriesPresenter() usecase.SearchRunnerHistoriesPresenter {
	return &searchRunnerHistoriesPresenter{}
}

func (p *searchRunnerHistoriesPresenter) Output(histories []domain.RunnerHistory) []usecase.SearchRunnerHistoriesOutput {
	result := make([]usecase.SearchRunnerHistoriesOutput, len(histories))
	for i, h := range histories {
		result[i] = usecase.SearchRunnerHistoriesOutput{
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
