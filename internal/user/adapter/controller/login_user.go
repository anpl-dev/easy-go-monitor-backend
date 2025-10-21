package controller

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginUserController struct {
	uc usecase.LoginUserUseCase
}

func NewLoginUserController(uc usecase.LoginUserUseCase) *LoginUserController {
	return &LoginUserController{uc: uc}
}

func (h *LoginUserController) Handle(c *gin.Context) {
	var input usecase.LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, err)
	}

	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, http.StatusOK, output)
}
