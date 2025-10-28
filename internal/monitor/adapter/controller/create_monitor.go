package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMonitorController struct {
	uc usecase.CreateMonitorUseCase
}

func NewCreateMonitorController(uc usecase.CreateMonitorUseCase) *CreateMonitorController {
	return &CreateMonitorController{uc: uc}
}

func (h *CreateMonitorController) Handle(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		response.HandleError(c, codes.ErrAuthFailed)
		return
	}

	userID, ok := userIDVal.(string)
	if !ok {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}

	var input usecase.CreateMonitorInput
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
