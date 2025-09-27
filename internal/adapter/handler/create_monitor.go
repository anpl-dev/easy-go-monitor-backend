package handler

import (
	"go-monitor-tool/internal/adapter/response"
	"go-monitor-tool/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMonitorHandler struct {
	uc usecase.CreateMonitorUseCase
}

func NewCreateMonitorHandler(uc usecase.CreateMonitorUseCase) *CreateMonitorHandler {
	return &CreateMonitorHandler{uc: uc}
}

func (h *CreateMonitorHandler) Handle(c *gin.Context) {
	var input usecase.CreateMonitorInput
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
