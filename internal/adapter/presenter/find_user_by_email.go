package presenter

import (
	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/usecase"
)

type FindUserByEmailPresenter struct{}

func NewFindUserByEmailPresenter() *FindUserByEmailPresenter {
	return &FindUserByEmailPresenter{}
}

func (p FindUserByEmailPresenter) Output(u *domain.User) usecase.FindUserByEmailOutput {
	if u == nil {
		return usecase.FindUserByEmailOutput{}
	}
	return usecase.FindUserByEmailOutput{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
