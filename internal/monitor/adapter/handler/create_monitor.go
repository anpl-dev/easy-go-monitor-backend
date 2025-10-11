package handler

import (
	"go-monitor-tool/internal/api/response"
	"go-monitor-tool/internal/apperr"
	"go-monitor-tool/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateMonitorHandler struct {
	uc usecase.CreateMonitorUseCase
}

func NewCreateMonitorHandler(uc usecase.CreateMonitorUseCase) *CreateMonitorHandler {
	return &CreateMonitorHandler{uc: uc}
}

func (h *CreateMonitorHandler) Handle(c *gin.Context) {
	userIDstr := c.GetString("user_id")
	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		response.HandleError(c, apperr.ErrInvalidUUID)
	}

	var input usecase.CreateMonitorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.HandleError(c, apperr.ErrBadRequest)
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
