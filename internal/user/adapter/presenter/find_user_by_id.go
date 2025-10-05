package presenter

import (
	"go-monitor-tool/internal/user/domain"
	"go-monitor-tool/internal/user/usecase"
)

type FindUserByIDPresenter struct{}

func NewFindUserByIDPresenter() *FindUserByIDPresenter {
	return &FindUserByIDPresenter{}
}

func (p *FindUserByIDPresenter) Output(user *domain.User) usecase.FindUserByIDOutput {
	if user == nil {
		return usecase.FindUserByIDOutput{}
	}
	return usecase.FindUserByIDOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
