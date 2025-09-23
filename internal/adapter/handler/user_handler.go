package handler

import (
	"go-monitor-tool/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	createUC usecase.CreateUserUseCase
}

// NewUserHandler constructor
func NewUserHandler(
	createUC usecase.CreateUserUseCase,
) *UserHandler {
	return &UserHandler{
		createUC: createUC,
	}
}

// CreateUser handles POST api /users
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
