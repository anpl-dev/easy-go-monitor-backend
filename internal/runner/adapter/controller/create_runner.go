package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/runner/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateRunnerController struct {
	uc usecase.CreateRunnerUseCase
}

func NewCreateRunnerController(uc usecase.CreateRunnerUseCase) *CreateRunnerController {
	return &CreateRunnerController{uc: uc}
}

func (h *CreateRunnerController) Handle(c *gin.Context) {
	userIDstr := c.GetString("user_id")
	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}

	var input usecase.CreateRunnerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, codes.ErrBadRequest)
		return
	}
	input.UserID = userID

	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.Success(c, http.StatusCreated, output)
}
