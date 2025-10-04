package handler

import (
	"go-monitor-tool/internal/adapter/response"
	"go-monitor-tool/internal/errors"
	"go-monitor-tool/internal/usecase"
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
		response.NewHTTPError(errors.ErrSearchParameters).Send(c)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.NewHTTPError(errors.ErrNotFound).Send(c)
		return
	}
	c.JSON(http.StatusOK, output)
}
