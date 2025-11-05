package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/infra/logger"
	"easy-go-monitor/internal/monitor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMonitorController struct {
	uc     usecase.CreateMonitorUseCase
	logger *logger.Logger
}

func NewCreateMonitorController(uc usecase.CreateMonitorUseCase, log *logger.Logger) *CreateMonitorController {
	return &CreateMonitorController{uc: uc, logger: log}
}

func (h *CreateMonitorController) Handle(c *gin.Context) {
	h.logger.Debug("Received CreateMonitor request", "path", c.FullPath())

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
		h.logger.Error("Invalid JSON input", "error", err)
		response.HandleError(c, codes.ErrBadRequest)
		return
	}

	input.UserID = userID

	output, err := h.uc.Execute(c.Request.Context(), input)
	if err != nil {
		h.logger.Error("CreateMonitor failed", "error", err)
		response.HandleError(c, err)
		return
	}
	h.logger.Debug("Monitor created successfully", "id", output.ID)
	response.Success(c, http.StatusCreated, output)
}
