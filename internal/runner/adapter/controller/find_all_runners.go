package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/runner/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FindAllRunnersController struct {
	uc usecase.FindAllRunnersUseCase
}

func NewFindAllRunnersController(uc usecase.FindAllRunnersUseCase) *FindAllRunnersController {
	return &FindAllRunnersController{uc: uc}
}

func (h *FindAllRunnersController) Handle(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		response.HandleError(c, codes.ErrAuthFailed)
		return
	}

	userIDStr, ok := userIDVal.(string)
	if !ok {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}

	output, err := h.uc.Execute(c.Request.Context(), usecase.FindAllRunnersInput{UserID: userID})
	if err != nil {
		response.HandleError(c, codes.ErrNotFound)
		return
	}
	response.Success(c, http.StatusOK, output)
}
