package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchMonitorsController struct {
	uc usecase.SearchMonitorsUseCase
}

func NewSearchMonitorsController(uc usecase.SearchMonitorsUseCase) *SearchMonitorsController {
	return &SearchMonitorsController{uc: uc}
}

func (h *SearchMonitorsController) Handle(c *gin.Context) {
	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.SearchMonitorsInput{UserID: userIDStr})
	if err != nil {
		response.HandleError(c, codes.ErrNotFound)
		return
	}
	response.Success(c, http.StatusOK, output)
}
