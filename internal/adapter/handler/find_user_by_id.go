package handler

import (
	"go-monitor-tool/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FindUserByIDHandler struct {
	findUC usecase.FindUserByIDUseCase
}

func NewFindUserByIDHandler(findUC usecase.FindUserByIDUseCase) *FindUserByIDHandler {
	return &FindUserByIDHandler{findUC: findUC}
}

func (h *FindUserByIDHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	output, err := h.findUC.Execute(c.Request.Context(), usecase.FindUserByIDInput{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
