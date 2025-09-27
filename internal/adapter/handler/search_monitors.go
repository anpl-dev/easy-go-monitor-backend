package handler

import (
	"go-monitor-tool/internal/adapter/response"
	"go-monitor-tool/internal/errors"
	"go-monitor-tool/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SearchMonitorsHandler struct {
	uc usecase.SearchMonitorsUseCase
}

func NewSearchMonitorsHandler(uc usecase.SearchMonitorsUseCase) *SearchMonitorsHandler {
	return &SearchMonitorsHandler{uc: uc}
}

func (h *SearchMonitorsHandler) Handle(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		response.NewHTTPError(errors.ErrInvalidUUID).Send(c)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.SearchMonitorsInput{UserID: userID})
	if err != nil {
		response.NewHTTPError(errors.ErrNotFound).Send(c)
		return
	}
	c.JSON(http.StatusOK, output)
}
