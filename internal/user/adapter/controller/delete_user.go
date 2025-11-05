package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteUserController struct {
	uc usecase.DeleteUserUseCase
}

func NewDeleteUserController(uc usecase.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{uc: uc}
}

func (h *DeleteUserController) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}

	err = h.uc.Execute(c.Request.Context(), usecase.DeleteUserInput{ID: id})
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, http.StatusNoContent, nil)
}
