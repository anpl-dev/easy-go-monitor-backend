package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/infra/logger"
	"easy-go-monitor/internal/runner/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ExecuteRunnerController struct {
	uc  usecase.ExecuteRunnerUseCase
	log *logger.Logger
}

func NewExecuteRunnerController(uc usecase.ExecuteRunnerUseCase, log *logger.Logger) *ExecuteRunnerController {
	return &ExecuteRunnerController{uc: uc, log: log}
}

func (h *ExecuteRunnerController) Handle(c *gin.Context) {
	idStr := c.Param("id")
	h.log.Debug("ExecuteRunner called", "runner_id", idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}

	output, err := h.uc.Execute(c.Request.Context(), usecase.ExecuteRunnerInput{
		RunnerIDs: []uuid.UUID{id},
	})
	if err != nil {
		h.log.Error("ExecuteRunner failed", "runner_id", id.String(), "error", err)
		response.HandleError(c, err)
		return
	}

	h.log.Debug("ExecuteRunner completed", "runner_id", id.String(), "result_count", len(output))
	response.Success(c, http.StatusOK, output)
}
