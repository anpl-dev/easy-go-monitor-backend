package handler

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/apperr"
	"go-monitor-tool/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchUsersHandler struct {
	uc usecase.SearchUsersUseCase
}

func NewSearchUsersHandler(uc usecase.SearchUsersUseCase) *SearchUsersHandler {
	return &SearchUsersHandler{uc: uc}
}

func (h *SearchUsersHandler) Handle(c *gin.Context) {
	input := usecase.SearchUsersInput{
		Email: c.Query("email"),
		Name:  c.Query("name"),
	}
	if input.Email == "" && input.Name == "" {
		response.HandleError(c, apperr.ErrSearchParameters)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, apperr.ErrNotFound)
		return
	}
	response.Success(c, http.StatusOK, output)
}
