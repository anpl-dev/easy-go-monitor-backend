package handler

import (
	"go-monitor-tool/internal/usecase"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid monitor id"})
		return
	}

	err = h.uc.Execute(c.Request.Context(), usecase.DeleteMonitorInput{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
