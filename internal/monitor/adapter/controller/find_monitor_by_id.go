package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FindMonitorByIDController struct {
	uc usecase.FindMonitorByIDUseCase
}

func NewFindMonitorByIDController(uc usecase.FindMonitorByIDUseCase) *FindMonitorByIDController {
	return &FindMonitorByIDController{uc: uc}
}

func (h *FindMonitorByIDController) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.FindMonitorByIDInput{ID: id})
	if err != nil {
		response.HandleError(c, codes.ErrNotFound)
		return
	}
	response.Success(c, http.StatusOK, output)
}
