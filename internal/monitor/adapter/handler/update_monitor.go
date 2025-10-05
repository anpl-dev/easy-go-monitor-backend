package handler

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/apperr"
	"go-monitor-tool/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateMonitorHandler struct {
	uc usecase.UpdateMonitorUseCase
}

func NewUpdateMonitorHandler(uc usecase.UpdateMonitorUseCase) *UpdateMonitorHandler {
	return &UpdateMonitorHandler{uc: uc}
}

func (h *UpdateMonitorHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, apperr.ErrInvalidUUID)
		return
	}

	var input usecase.UpdateMonitorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, apperr.ErrBadRequest)
		return
	}

	input.ID = id

	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, http.StatusOK, output)
}
