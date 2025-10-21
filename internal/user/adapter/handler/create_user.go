package handler

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/constraints"
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
		response.HandleError(c, constraints.ErrBadRequest)
		return
	}

	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, http.StatusCreated, output)
}
