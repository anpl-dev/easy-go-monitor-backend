package presenter

import (
	"go-monitor-tool/internal/monitor/usecase"
	"go-monitor-tool/internal/monitor/domain"
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
		ID:             monitor.ID,
		UserID:         monitor.UserID,
		Name:           monitor.Name,
		URL:            monitor.URL,
		IntervalSecond: monitor.IntervalSecond,
		CreatedAt:      monitor.CreatedAt,
		UpdatedAt:      monitor.UpdatedAt,
	}
}
