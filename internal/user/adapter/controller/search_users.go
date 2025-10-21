package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/constraint"
	"easy-go-monitor/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchUsersController struct {
	uc usecase.SearchUsersUseCase
}

func NewSearchUsersController(uc usecase.SearchUsersUseCase) *SearchUsersController {
	return &SearchUsersController{uc: uc}
}

func (h *SearchUsersController) Handle(c *gin.Context) {
	input := usecase.SearchUsersInput{
		Email: c.Query("email"),
		Name:  c.Query("name"),
	}
	if input.Email == "" && input.Name == "" {
		response.HandleError(c, constraint.ErrSearchParameters)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, constraint.ErrNotFound)
		return
	}
	response.Success(c, http.StatusOK, output)
}
