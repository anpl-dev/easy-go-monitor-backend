package handler

import (
	"go-monitor-tool/internal/response"
	"go-monitor-tool/internal/errors"
	"go-monitor-tool/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteUserHandler struct {
	uc usecase.DeleteUserUseCase
}

func NewDeleteUserHandler(uc usecase.DeleteUserUseCase) *DeleteUserHandler {
	return &DeleteUserHandler{uc: uc}
}

func (h *DeleteUserHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.NewHTTPError(errors.ErrInvalidUUID).Send(c)
		return
	}

	err = h.uc.Execute(c.Request.Context(), usecase.DeleteUserInput{ID: id})
	if err != nil {
		response.NewHTTPError(err).Send(c)
		return
	}
	c.Status(http.StatusNoContent)
}
