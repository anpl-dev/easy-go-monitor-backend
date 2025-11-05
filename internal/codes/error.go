package codes

import "fmt"

type Error struct {
	Code    int
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e *Error) Unwrap() error {
	return e.Err
}

func Wrap(base *Error, err error) *Error {
	return &Error{
		Code:    base.Code,
		Message: base.Message,
		Err:     err,
	}
}
