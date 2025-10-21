package controller

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/constraints"
	"go-monitor-tool/internal/user/usecase"
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
		response.HandleError(c, constraints.ErrSearchParameters)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, constraints.ErrNotFound)
		return
	}
	response.Success(c, http.StatusOK, output)
}
