package handler

import (
	"go-monitor-tool/internal/adapter/response"
	"go-monitor-tool/internal/errors"
	"go-monitor-tool/internal/usecase"
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
		response.NewHTTPError(errors.ErrInvalidUUID).Send(c)
		return
	}

	var input usecase.UpdateMonitorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(c)
		return
	}

	input.ID = id

	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.NewHTTPError(err).Send(c)
		return
	}
	c.JSON(http.StatusOK, output)
}
