package presenter

import (
	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/usecase"
)

type FindUserByIDPresenter struct{}

func NewFindUserByIDPresenter() *FindUserByIDPresenter {
	return &FindUserByIDPresenter{}
}

func (p FindUserByIDPresenter) Output(u *domain.User) usecase.FindUserByIDOutput {
	if u == nil {
		return usecase.FindUserByIDOutput{}
	}
	return usecase.FindUserByIDOutput{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
