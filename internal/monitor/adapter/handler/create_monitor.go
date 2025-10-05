package handler

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/apperr"
	"go-monitor-tool/internal/monitor/usecase"
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
		response.HandleError(c, apperr.ErrBadRequest)
		return
	}

	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, http.StatusCreated, output)
}
