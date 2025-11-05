package presenter

import (
	"easy-go-monitor/internal/monitor/domain"
	"easy-go-monitor/internal/monitor/usecase"
)

type SetEnabledMonitorPresenter struct{}

func NewSetEnabledMonitorPresenter() *SetEnabledMonitorPresenter {
	return &SetEnabledMonitorPresenter{}
}

func (p *SetEnabledMonitorPresenter) Output(monitor *domain.Monitor) usecase.SetEnabledMonitorOutput {
	if monitor == nil {
		return usecase.SetEnabledMonitorOutput{}
	}
	return usecase.SetEnabledMonitorOutput{
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
