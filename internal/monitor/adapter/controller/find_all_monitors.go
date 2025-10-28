package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindAllMonitorsController struct {
	uc usecase.FindAllMonitorsUseCase
}

func NewFindAllMonitorsController(uc usecase.FindAllMonitorsUseCase) *FindAllMonitorsController {
	return &FindAllMonitorsController{uc: uc}
}

func (h *FindAllMonitorsController) Handle(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		response.HandleError(c, codes.ErrAuthFailed)
		return
	}

	userID, ok := userIDVal.(string)
	if !ok {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.FindAllMonitorsInput{UserID: userID})
	if err != nil {
		response.HandleError(c, codes.ErrNotFound)
		return
	}
	response.Success(c, http.StatusOK, output)
}
