package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SetEnabledMonitorController struct {
	uc usecase.SetEnabledMonitorUseCase
}

func NewSetEnabledMonitorController(uc usecase.SetEnabledMonitorUseCase) *SetEnabledMonitorController {
	return &SetEnabledMonitorController{uc: uc}
}

func (h *SetEnabledMonitorController) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}

	var input usecase.SetEnabledMonitorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, codes.ErrBadRequest)
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
