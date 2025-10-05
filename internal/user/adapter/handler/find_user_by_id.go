package handler

import (
	"go-monitor-tool/internal/response"
	"go-monitor-tool/internal/errors"
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
		response.NewHTTPError(errors.ErrInvalidUUID).Send(c)
		return
	}
	output, err := h.uc.Execute(c.Request.Context(), usecase.FindUserByIDInput{ID: id})
	if err != nil {
		response.NewHTTPError(errors.ErrNotFound).Send(c)
		return
	}
	c.JSON(http.StatusOK, output)
}
