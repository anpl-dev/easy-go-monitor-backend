package handler

import (
	"go-monitor-tool/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MonitorHandler handles HTTP requests for monitors.
type MonitorHandler struct {
	createUC usecase.CreateMonitorUseCase
}

// NewMonitorHandler constructor
func NewMonitorHandler(createUC usecase.CreateMonitorUseCase) *MonitorHandler {
	return &MonitorHandler{createUC: createUC}
}

// CreateMonitor handles POST /monitors
func (h *MonitorHandler) CreateMonitor(c *gin.Context) {
	var input usecase.CreateMonitorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := h.createUC.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, output)
}
