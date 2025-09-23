package presenter

import (
	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/usecase"
)

type UserHTTPPresenter struct{}

func (p UserHTTPPresenter) Output(u *domain.User) usecase.CreateUserOutput {
	return usecase.CreateUserOutput{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
