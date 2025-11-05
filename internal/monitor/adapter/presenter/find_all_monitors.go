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
	for _, monitor := range monitors {
		if monitor == nil {
			continue
		}
		outputs = append(outputs, usecase.FindAllMonitorsOutput{
			ID:        monitor.ID,
			UserID:    monitor.UserID,
			Name:      monitor.Name,
			URL:       monitor.URL,
			Type:      monitor.Type,
			Settings:  monitor.Settings,
			IsEnabled: monitor.IsEnabled,
			CreatedAt: monitor.CreatedAt,
			UpdatedAt: monitor.UpdatedAt,
		})
	}
	return outputs
}
