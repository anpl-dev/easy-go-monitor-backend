package handler

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/apperr"
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
		response.HandleError(c, apperr.ErrInvalidUUID)
		return
	}

	err = h.uc.Execute(c.Request.Context(), usecase.DeleteUserInput{ID: id})
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, http.StatusNoContent, nil)
}
