package handler

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/apperr"
	"go-monitor-tool/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteMonitorHandler struct {
	uc usecase.DeleteMonitorUseCase
}

func NewDeleteMonitorHandler(uc usecase.DeleteMonitorUseCase) *DeleteMonitorHandler {
	return &DeleteMonitorHandler{uc: uc}
}

func (h *DeleteMonitorHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, apperr.ErrInvalidUUID)
		return
	}

	err = h.uc.Execute(c.Request.Context(), usecase.DeleteMonitorInput{ID: id})
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, http.StatusNoContent, nil)
}
