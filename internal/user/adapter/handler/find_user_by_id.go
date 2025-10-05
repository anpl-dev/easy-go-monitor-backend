package handler

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/apperr"
	"go-monitor-tool/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FindUserByIDHandler struct {
	uc usecase.FindUserByIDUseCase
}

func NewFindUserByIDHandler(uc usecase.FindUserByIDUseCase) *FindUserByIDHandler {
	return &FindUserByIDHandler{uc: uc}
}

func (h *FindUserByIDHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, apperr.ErrInvalidUUID)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.FindUserByIDInput{ID: id})
	if err != nil {
		response.HandleError(c, apperr.ErrNotFound)
		return
	}
	response.Success(c, http.StatusOK, output)
}
