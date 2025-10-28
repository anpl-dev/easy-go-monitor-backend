package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/runner/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FindRunnerByIDController struct {
	uc usecase.FindRunnerByIDUseCase
}

func NewFindRunnerByIDController(uc usecase.FindRunnerByIDUseCase) *FindRunnerByIDController {
	return &FindRunnerByIDController{uc: uc}
}

func (h *FindRunnerByIDController) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.FindRunnerByIDInput{ID: id})
	if err != nil {
		response.HandleError(c, codes.ErrNotFound)
		return
	}
	response.Success(c, http.StatusOK, output)
}
