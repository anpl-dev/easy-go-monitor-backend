package apperr

import "fmt"

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func New(code int, message string) *AppError {
	return &AppError{
		Code: code, 
		Message: message,
	}
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func Wrap(base *AppError, err error) *AppError {
	return &AppError{
		Code:    base.Code,
		Message: base.Message,
		Err:     err,
	}
}
