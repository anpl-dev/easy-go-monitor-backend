package presenter

import (
	"easy-go-monitor/internal/monitor/domain"
	"easy-go-monitor/internal/monitor/usecase"
)

type FindMonitorByIDPresenter struct{}

func NewFindMonitorByIDPresenter() *FindMonitorByIDPresenter {
	return &FindMonitorByIDPresenter{}
}

func (p *FindMonitorByIDPresenter) Output(monitor *domain.Monitor) usecase.FindMonitorByIDOutput {
	if monitor == nil {
		return usecase.FindMonitorByIDOutput{}
	}
	return usecase.FindMonitorByIDOutput{
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
