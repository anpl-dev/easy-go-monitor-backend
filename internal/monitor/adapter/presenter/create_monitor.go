package presenter

import (
	"easy-go-monitor/internal/monitor/domain"
	"easy-go-monitor/internal/monitor/usecase"
)

type CreateMonitorPresenter struct{}

func NewCreateMonitorPresenter() *CreateMonitorPresenter {
	return &CreateMonitorPresenter{}
}

func (p *CreateMonitorPresenter) Output(monitor *domain.Monitor) usecase.CreateMonitorOutput {
	if monitor == nil {
		return usecase.CreateMonitorOutput{}
	}
	return usecase.CreateMonitorOutput{
		ID:        monitor.ID,
		UserID:    monitor.UserID,
		Name:      monitor.Name,
		URL:       monitor.URL,
		Type:      monitor.Type,
		Settings:  monitor.Settings,
		IsEnabled: monitor.IsEnabled,
		CreatedAt: monitor.CreatedAt,
		UpdatedAt: monitor.UpdatedAt,
	}
}
