package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteMonitorController struct {
	uc usecase.DeleteMonitorUseCase
}

func NewDeleteMonitorController(uc usecase.DeleteMonitorUseCase) *DeleteMonitorController {
	return &DeleteMonitorController{uc: uc}
}

func (h *DeleteMonitorController) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}

	err = h.uc.Execute(c.Request.Context(), usecase.DeleteMonitorInput{ID: id})
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, http.StatusNoContent, nil)
}
