package handler

import (
	"go-monitor-tool/internal/response"
	"go-monitor-tool/internal/errors"
	"go-monitor-tool/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateUserHandler struct {
	uc usecase.UpdateUserUseCase
}

func NewUpdateUserHandler(uc usecase.UpdateUserUseCase) *UpdateUserHandler {
	return &UpdateUserHandler{uc: uc}
}

func (h *UpdateUserHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.NewHTTPError(errors.ErrInvalidUUID).Send(c)
		return
	}

	var input usecase.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.NewError(http.StatusBadRequest, err).Send(c)
		return
	}

	input.ID = id

	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.NewHTTPError(err).Send(c)
		return
	}
	c.JSON(http.StatusOK, output)
}
