package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/constraint"
	"easy-go-monitor/internal/user/usecase"
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
		response.HandleError(c, constraint.ErrInvalidUUID)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.FindUserByIDInput{ID: id})
	if err != nil {
		response.HandleError(c, constraint.ErrNotFound)
		return
	}
	response.Success(c, http.StatusOK, output)
}
