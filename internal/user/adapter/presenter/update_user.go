package presenter

import (
	"go-monitor-tool/internal/user/domain"
	"go-monitor-tool/internal/user/usecase"
)

type UpdateUserPresenter struct{}

func NewUpdateUserPresenter() *UpdateUserPresenter {
	return &UpdateUserPresenter{}
}

func (p *UpdateUserPresenter) Output(user *domain.User) usecase.UpdateUserOutput {
	if user == nil {
		return usecase.UpdateUserOutput{}
	}
	return usecase.UpdateUserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}
}
