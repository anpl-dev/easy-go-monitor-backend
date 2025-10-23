package response

import (
	"easy-go-monitor/internal/codes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func ErrorResponse(err error) (int, APIResponse) {
	if res, ok := err.(*codes.Error); ok {
		status := toHTTPStatus(res.Code)
		return status, APIResponse{
			Code:    res.Code,
			Message: res.Message,
			Error:   res.Error(),
		}
	}

	return http.StatusInternalServerError, APIResponse{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
		Error:   err.Error(),
	}
}

func HandleError(c *gin.Context, err error) {
	status, res := ErrorResponse(err)
	c.JSON(status, res)
	c.Abort()
}

func toHTTPStatus(code int) int {
	switch {
	case code >= 4000 && code < 4100:
		return http.StatusBadRequest
	case code >= 4100 && code < 4200:
		return http.StatusUnauthorized
	case code >= 5000:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
