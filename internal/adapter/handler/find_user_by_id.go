package handler

import (
	"go-monitor-tool/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FindUserByIDHandler struct {
	uc usecase.FindUserByIDUseCase
}

func NewFindUserByIDHandler(findUC usecase.FindUserByIDUseCase) *FindUserByIDHandler {
	return &FindUserByIDHandler{uc: findUC}
}

func (h *FindUserByIDHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.FindUserByIDInput{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
