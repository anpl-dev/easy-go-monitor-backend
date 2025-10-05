package presenter

import (
	"go-monitor-tool/internal/user/domain"
	"go-monitor-tool/internal/user/usecase"
)

type SearchUsersPresenter struct{}

func NewSearchUsersPresenter() *SearchUsersPresenter {
	return &SearchUsersPresenter{}
}

func (p *SearchUsersPresenter) Output(users []*domain.User) []usecase.SearchUsersOutput {
	output := make([]usecase.SearchUsersOutput, 0, len(users))
	for _, user := range users {
		output = append(output, usecase.SearchUsersOutput{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return output
}
