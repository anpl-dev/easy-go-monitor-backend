package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/runner/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FindRunnerHistoriesController struct {
	uc usecase.FindRunnerHistoriesUseCase
}

func NewFindRunnerHistoriesController(uc usecase.FindRunnerHistoriesUseCase) *FindRunnerHistoriesController {
	return &FindRunnerHistoriesController{uc: uc}
}

func (h *FindRunnerHistoriesController) Handle(c *gin.Context) {
	idStr := c.Param("id")
	runnerID, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}

	output, err := h.uc.Execute(c.Request.Context(), usecase.FindRunnerHistoriesInput{RunnerID: runnerID})
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, http.StatusOK, output)
}
