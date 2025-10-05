package handler

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginUserHandler struct {
	uc usecase.LoginUserUseCase
}

func NewLoginUserHandler(uc usecase.LoginUserUseCase) *LoginUserHandler {
	return &LoginUserHandler{uc: uc}
}

func (h *LoginUserHandler) Handle(c *gin.Context) {
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
