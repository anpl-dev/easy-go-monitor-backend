package presenter

import (
	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/usecase"
)

type MonitorHTTPPresenter struct{}

func (p MonitorHTTPPresenter) Output(m *domain.Monitor) usecase.CreateMonitorOutput {
    return usecase.CreateMonitorOutput{
        ID:             m.ID,
        UserID:         m.UserID,
        Name:           m.Name,
        URL:            m.Url,
        IntervalSecond: m.IntervalSeconds,
        CreatedAt:      m.CreatedAt,
        UpdatedAt:      m.UpdatedAt,
    }
}