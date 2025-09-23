package presenter

import (
	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/usecase"
)

type CreateMonitorPresenter struct{}

func NewCreateMonitorPresenter() *CreateMonitorPresenter {
    return &CreateMonitorPresenter{}
}

func (p CreateMonitorPresenter) Output(m *domain.Monitor) usecase.CreateMonitorOutput {
    return usecase.CreateMonitorOutput{
        ID:             m.ID,
        UserID:         m.UserID,
        Name:           m.Name,
        Url:            m.Url,
        IntervalSecond: m.IntervalSecond,
        CreatedAt:      m.CreatedAt,
        UpdatedAt:      m.UpdatedAt,
    }
}