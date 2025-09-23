package presenter

import (
	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/usecase"
)

type FindMonitorsByUserPresenter struct{}

func NewFindMonitorsByUserPresenter() *FindMonitorsByUserPresenter {
	return &FindMonitorsByUserPresenter{}
}

func (p *FindMonitorsByUserPresenter) Output(monitors []*domain.Monitor) []usecase.FindMonitorsByUserOutput {
	if monitors == nil {
		return []usecase.FindMonitorsByUserOutput{}
	}

	outputs := make([]usecase.FindMonitorsByUserOutput, 0, len(monitors))
	for _, m := range monitors {
		if m == nil {
			continue
		}
		outputs = append(outputs, usecase.FindMonitorsByUserOutput{
			ID:             m.ID,
			UserID:         m.UserID,
			Name:           m.Name,
			Url:            m.Url,
			IntervalSecond: m.IntervalSecond,
			CreatedAt:      m.CreatedAt,
			UpdatedAt:      m.UpdatedAt,
		})
	}
	return outputs
}
