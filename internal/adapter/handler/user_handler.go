package handler

import (
	"go-monitor-tool/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createUC usecase.CreateUserUseCase
}

func NewCreateUser(createUC usecase.CreateUserUseCase) *UserHandler {
	return &UserHandler{createUC: createUC}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input usecase.CreateUserInput
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
