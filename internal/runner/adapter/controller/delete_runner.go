package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/runner/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteRunnerController struct {
	uc usecase.DeleteRunnerUseCase
}

func NewDeleteRunnerController(uc usecase.DeleteRunnerUseCase) *DeleteRunnerController {
	return &DeleteRunnerController{uc: uc}
}

func (h *DeleteRunnerController) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}

	err = h.uc.Execute(c.Request.Context(), usecase.DeleteRunnerInput{ID: id})
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, http.StatusNoContent, nil)
}
