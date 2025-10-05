package presenter

import (
	"go-monitor-tool/internal/user/domain"
	"go-monitor-tool/internal/user/usecase"
)

type CreateUserPresenter struct{}

func NewCreateUserPresenter() *CreateUserPresenter {
	return &CreateUserPresenter{}
}

func (p *CreateUserPresenter) Output(user *domain.User) usecase.CreateUserOutput {
	if user == nil {
		return usecase.CreateUserOutput{}
	}
	return usecase.CreateUserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
