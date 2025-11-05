package controller

import (
	"easy-go-monitor/internal/api/response"
	"easy-go-monitor/internal/codes"
	"easy-go-monitor/internal/runner/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SearchRunnerHistoriesController struct {
	uc usecase.SearchRunnerHistoriesUseCase
}

func NewSearchRunnerHistoriesController(uc usecase.SearchRunnerHistoriesUseCase) *SearchRunnerHistoriesController {
	return &SearchRunnerHistoriesController{uc: uc}
}

func (h *SearchRunnerHistoriesController) Handle(c *gin.Context) {
	userIDstr := c.GetString("user_id")
	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		response.HandleError(c, codes.ErrInvalidUUID)
		return
	}

	var input usecase.SearchRunnerHistoriesInput
	if err := c.ShouldBindQuery(&input); err != nil {
		response.HandleError(c, codes.ErrBadRequest)
		return
	}
	input.UserID = userID

	output, err := h.uc.Execute(c.Request.Context(), usecase.SearchRunnerHistoriesInput{
		UserID:  userID,
		Status:  input.Status,
		Minutes: input.Minutes,
	})
	if err != nil {
		response.HandleError(c, err)
		return
	}

	response.Success(c, http.StatusOK, output)
}
