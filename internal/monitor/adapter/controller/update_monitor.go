package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/constraints"
	"easy-go-monitor/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateMonitorController struct {
	uc usecase.UpdateMonitorUseCase
}

func NewUpdateMonitorController(uc usecase.UpdateMonitorUseCase) *UpdateMonitorController {
	return &UpdateMonitorController{uc: uc}
}

func (h *UpdateMonitorController) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, constraints.ErrInvalidUUID)
		return
	}

	var input usecase.UpdateMonitorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, constraints.ErrBadRequest)
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
