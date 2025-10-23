package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	uc usecase.CreateUserUseCase
}

func NewCreateUserController(uc usecase.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{uc: uc}
}

func (h *CreateUserController) Handle(c *gin.Context) {
	var input usecase.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, codes.ErrBadRequest)
		return
	}

	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, http.StatusCreated, output)
}
