package response

import (
	"go-monitor-tool/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	StatusCode int      `json:"code"`
	Errors     []string `json:"errors"`
}

func NewError(err error, status int) *Error {
	return &Error{
		StatusCode: status,
		Errors:     []string{err.Error()},
	}
}

func (e Error) Send(c *gin.Context) {
	c.JSON(e.StatusCode, e)
}

func NewHTTPError(err error) *Error {
	switch {
	// Domain
	case errors.Is(err, errors.ErrInvalidUserName),
		errors.Is(err, errors.ErrInvalidEmail),
		errors.Is(err, errors.ErrInvalidPassword):
		return NewError(err, http.StatusBadRequest)

	case errors.Is(err, errors.ErrInvalidMonitorName),
		errors.Is(err, errors.ErrInvalidMonitorURL),
		errors.Is(err, errors.ErrInvalidMonitorInterval):
		return NewError(err, http.StatusBadRequest)

	case errors.Is(err, errors.ErrNotFound):
		return NewError(err, http.StatusNotFound)

	case errors.Is(err, errors.ErrInvalidUUID):
		return NewError(err, http.StatusBadRequest)

	// Handler
	case errors.Is(err, errors.ErrSearchParameters):
		return NewError(err, http.StatusBadRequest)

	default:
		return NewError(errors.New("internal server error"), http.StatusInternalServerError)
	}
}
