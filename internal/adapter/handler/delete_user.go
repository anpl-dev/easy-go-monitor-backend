package handler

import (
	"go-monitor-tool/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteUserHandler struct {
	uc usecase.DeleteUserUseCase
}

func NewDeleteUserHandler(uc usecase.DeleteUserUseCase) *DeleteUserHandler {
	return &DeleteUserHandler{uc: uc}
}

func (h *DeleteUserHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	err = h.uc.Execute(c.Request.Context(), usecase.DeleteUserInput{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
