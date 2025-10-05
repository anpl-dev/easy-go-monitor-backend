package handler

import (
	"go-monitor-tool/internal/response"
	"go-monitor-tool/internal/errors"
	"go-monitor-tool/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FindMonitorByIDHandler struct {
	uc usecase.FindMonitorByIDUseCase
}

func NewFindMonitorByIDHandler(uc usecase.FindMonitorByIDUseCase) *FindMonitorByIDHandler {
	return &FindMonitorByIDHandler{uc: uc}
}

func (h *FindMonitorByIDHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.NewHTTPError(errors.ErrInvalidUUID).Send(c)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.FindMonitorByIDInput{ID: id})
	if err != nil {
		response.NewHTTPError(errors.ErrNotFound).Send(c)
		return
	}
	c.JSON(http.StatusOK, output)
}
