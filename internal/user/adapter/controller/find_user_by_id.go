package controller

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/constraints"
	"go-monitor-tool/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FindUserByIDController struct {
	uc usecase.FindUserByIDUseCase
}

func NewFindUserByIDController(uc usecase.FindUserByIDUseCase) *FindUserByIDController {
	return &FindUserByIDController{uc: uc}
}

func (h *FindUserByIDController) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, constraints.ErrInvalidUUID)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.FindUserByIDInput{ID: id})
	if err != nil {
		response.HandleError(c, constraints.ErrNotFound)
		return
	}
	response.Success(c, http.StatusOK, output)
}
