package presenter

import (
	"easy-go-monitor/internal/monitor/domain"
	"easy-go-monitor/internal/monitor/usecase"
)

type SearchMonitorsPresenter struct{}

func NewSearchMonitorsPresenter() *SearchMonitorsPresenter {
	return &SearchMonitorsPresenter{}
}

func (p *SearchMonitorsPresenter) Output(monitors []*domain.Monitor) []usecase.SearchMonitorsOutput {
	if monitors == nil {
		return []usecase.SearchMonitorsOutput{}
	}

	outputs := make([]usecase.SearchMonitorsOutput, 0, len(monitors))
	for _, m := range monitors {
		if m == nil {
			continue
		}
		outputs = append(outputs, usecase.SearchMonitorsOutput{
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
