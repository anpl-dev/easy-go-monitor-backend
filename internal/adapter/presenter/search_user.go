package presenter

import (
	"go-monitor-tool/internal/domain"
	"go-monitor-tool/internal/usecase"
)

type SearchUserPresenter struct{}

func NewSearchUserPresenter() *SearchUserPresenter {
	return &SearchUserPresenter{}
}

func (p *SearchUserPresenter) Output(users []*domain.User) []usecase.SearchUserOutput {
	output := make([]usecase.SearchUserOutput, 0, len(users))
	for _, user := range users {
		output = append(output, usecase.SearchUserOutput{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return output
}
