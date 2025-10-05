package handler

import (
	"go-monitor-tool/internal/response"
	"go-monitor-tool/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserHandler struct {
	uc usecase.CreateUserUseCase
}

func NewCreateUserHandler(uc usecase.CreateUserUseCase) *CreateUserHandler {
	return &CreateUserHandler{uc: uc}
}

func (h *CreateUserHandler) Handle(c *gin.Context) {
	var input usecase.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.NewError(http.StatusBadRequest, err).Send(c)
		return
	}

	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.NewHTTPError(err).Send(c)
		return
	}
	c.JSON(http.StatusCreated, output)

}
