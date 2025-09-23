package handler

import (
	"go-monitor-tool/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindUserByEmailHandler struct {
	uc usecase.FindUserByEmailUseCase
}

func NewFindUserByEmailHandler(findUC usecase.FindUserByEmailUseCase) *FindUserByEmailHandler {
	return &FindUserByEmailHandler{uc: findUC}
}

func (h *FindUserByEmailHandler) Handle(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.FindUserByEmailInput{Email: email})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
