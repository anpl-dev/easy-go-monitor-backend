package handler

import (
	"go-monitor-tool/internal/adapter/response"
	"go-monitor-tool/internal/errors"
	"go-monitor-tool/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchUserHandler struct {
	uc usecase.SearchUserUseCase
}

func NewSearchUserHandler(uc usecase.SearchUserUseCase) *SearchUserHandler {
	return &SearchUserHandler{uc: uc}
}

func (h *SearchUserHandler) Handle(c *gin.Context) {
	input := usecase.SearchUserInput{
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
