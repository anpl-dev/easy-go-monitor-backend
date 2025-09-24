package handler

import (
	"go-monitor-tool/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FindMonitorsByUserHandler struct {
	uc usecase.FindMonitorsByUserUseCase
}

func NewFindMonitorsByUserHandler(uc usecase.FindMonitorsByUserUseCase) *FindMonitorsByUserHandler {
	return &FindMonitorsByUserHandler{uc: uc}
}

func (h *FindMonitorsByUserHandler) Handle(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.FindMonitorsByUserInput{UserID: userID})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
