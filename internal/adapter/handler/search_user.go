package handler

import (
	"go-monitor-tool/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchUserHandler struct {
	uc usecase.SearchUserUseCase
}

func NewSearchUserHandler(SearchUC usecase.SearchUserUseCase) *SearchUserHandler {
	return &SearchUserHandler{uc: SearchUC}
}

func (h *SearchUserHandler) Handle(c *gin.Context) {
	input := usecase.SearchUserInput{
		Email: c.Query("email"),
		Name:  c.Query("name"),
	}
	if input.Email == "" && input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no search parameters"})
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
