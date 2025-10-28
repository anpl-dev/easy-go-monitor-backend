package presenter

import (
	"easy-go-monitor/internal/monitor/domain"
	"easy-go-monitor/internal/monitor/usecase"
)

type FindAllMonitorsPresenter struct{}

func NewFindAllMonitorsPresenter() *FindAllMonitorsPresenter {
	return &FindAllMonitorsPresenter{}
}

func (p *FindAllMonitorsPresenter) Output(monitors []*domain.Monitor) []usecase.FindAllMonitorsOutput {
	if monitors == nil {
		return []usecase.FindAllMonitorsOutput{}
	}

	outputs := make([]usecase.FindAllMonitorsOutput, 0, len(monitors))
	for _, m := range monitors {
		if m == nil {
			continue
		}
		outputs = append(outputs, usecase.FindAllMonitorsOutput{
			ID:        m.ID,
			UserID:    m.UserID,
			Name:      m.Name,
			URL:       m.URL,
			Type:      m.Type,
			Settings:  m.Settings,
			IsEnabled: m.IsEnabled,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		})
	}
	return outputs
}
