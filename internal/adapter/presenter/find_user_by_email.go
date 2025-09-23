package presenter

import (
	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/usecase"
)

type FindUserByEmailPresenter struct{}

func NewFindUserByEmailPresenter() *FindUserByEmailPresenter {
	return &FindUserByEmailPresenter{}
}

func (p *FindUserByEmailPresenter) Output(user *domain.User) usecase.FindUserByEmailOutput {
	if user == nil {
		return usecase.FindUserByEmailOutput{}
	}
	return usecase.FindUserByEmailOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
