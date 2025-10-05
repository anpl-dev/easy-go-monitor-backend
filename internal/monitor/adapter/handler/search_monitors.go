package handler

import (
	"go-monitor-tool/internal/response"
	"go-monitor-tool/internal/errors"
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
		response.NewHTTPError(errors.ErrInvalidUUID).Send(c)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.SearchMonitorsInput{UserID: userIDStr})
	if err != nil {
		response.NewHTTPError(errors.ErrNotFound).Send(c)
		return
	}
	c.JSON(http.StatusOK, output)
}
