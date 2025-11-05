package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateUserController struct {
	uc usecase.UpdateUserUseCase
}

func NewUpdateUserController(uc usecase.UpdateUserUseCase) *UpdateUserController {
	return &UpdateUserController{uc: uc}
}

func (h *UpdateUserController) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}

	var input usecase.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, codes.ErrBadRequest)
		return
	}

	input.ID = id

	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, http.StatusOK, output)
}
