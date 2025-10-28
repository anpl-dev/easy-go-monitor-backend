package presenter

import (
	"easy-go-monitor/internal/monitor/domain"
	"easy-go-monitor/internal/monitor/usecase"
)

type UpdateMonitorPresenter struct{}

func NewUpdateMonitorPresenter() *UpdateMonitorPresenter {
	return &UpdateMonitorPresenter{}
}

func (p *UpdateMonitorPresenter) Output(monitor *domain.Monitor) usecase.UpdateMonitorOutput {
	if monitor == nil {
		return usecase.UpdateMonitorOutput{}
	}
	return usecase.UpdateMonitorOutput{
		ID:        monitor.ID,
		UserID:    monitor.UserID,
		Name:      monitor.Name,
		URL:       monitor.URL,
		Type:      monitor.Type,
		Settings:  monitor.Settings,
		IsEnabled: monitor.IsEnabled,
		UpdatedAt: monitor.UpdatedAt,
	}
}
