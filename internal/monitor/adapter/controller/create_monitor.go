package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/constraint"
	"easy-go-monitor/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateMonitorController struct {
	uc usecase.CreateMonitorUseCase
}

func NewCreateMonitorController(uc usecase.CreateMonitorUseCase) *CreateMonitorController {
	return &CreateMonitorController{uc: uc}
}

func (h *CreateMonitorController) Handle(c *gin.Context) {
	userIDstr := c.GetString("user_id")
	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		response.HandleError(c, constraint.ErrInvalidUUID)
	}

	var input usecase.CreateMonitorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, constraint.ErrBadRequest)
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
