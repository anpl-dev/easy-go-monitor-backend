package handler

import (
	"go-monitor-tool/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMonitorHandler struct {
	createUC usecase.CreateMonitorUseCase
}

func NewCreateMonitorHandler(createUC usecase.CreateMonitorUseCase) *CreateMonitorHandler {
	return &CreateMonitorHandler{createUC: createUC}
}

func (h *CreateMonitorHandler) Handle(c *gin.Context) {
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
