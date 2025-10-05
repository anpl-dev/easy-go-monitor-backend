package handler

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/apperr"
	"go-monitor-tool/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchMonitorsHandler struct {
	uc usecase.SearchMonitorsUseCase
}

func NewSearchMonitorsHandler(uc usecase.SearchMonitorsUseCase) *SearchMonitorsHandler {
	return &SearchMonitorsHandler{uc: uc}
}

func (h *SearchMonitorsHandler) Handle(c *gin.Context) {
	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		response.HandleError(c, apperr.ErrInvalidUUID)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.SearchMonitorsInput{UserID: userIDStr})
	if err != nil {
		response.HandleError(c, apperr.ErrNotFound)
		return
	}
	response.Success(c, http.StatusOK, output)
}
