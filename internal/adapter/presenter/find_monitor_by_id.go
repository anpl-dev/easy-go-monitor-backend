package presenter

import (
	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/usecase"
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
		ID:             monitor.ID,
		UserID:         monitor.UserID,
		Name:           monitor.Name,
		IntervalSecond: monitor.IntervalSecond,
		CreatedAt:      monitor.CreatedAt,
		UpdatedAt:      monitor.UpdatedAt,
	}
}
