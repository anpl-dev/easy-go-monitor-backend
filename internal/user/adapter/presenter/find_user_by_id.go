package presenter

import (
	"easy-go-monitor/internal/user/domain"
	"easy-go-monitor/internal/user/usecase"
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
