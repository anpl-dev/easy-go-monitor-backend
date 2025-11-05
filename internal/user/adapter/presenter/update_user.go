package presenter

import (
	"easy-go-monitor/internal/user/domain"
	"easy-go-monitor/internal/user/usecase"
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
