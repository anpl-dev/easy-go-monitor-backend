package presenter

import (
	"go-monitor-tool/internal/monitor/domain"
	"go-monitor-tool/internal/monitor/usecase"
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
		ID:             monitor.ID,
		UserID:         monitor.UserID,
		Name:           monitor.Name,
		URL:            monitor.URL,
		IntervalSecond: monitor.IntervalSecond,
		UpdatedAt:      monitor.UpdatedAt,
	}
}
