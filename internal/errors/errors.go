package errors

import (
	"errors"
)

// Domain
var (
	// Common errors
	ErrNotFound           = errors.New("not found")
	ErrInvalidUUID        = errors.New("invalid uuid")
	ErrInvalidCredentials = errors.New("invalid credentials")

	// User errors
	ErrInvalidUserName = errors.New("invalid user name")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password hash")

	// Monitor errors
	ErrInvalidMonitorName     = errors.New("invalid monitor name")
	ErrInvalidMonitorURL      = errors.New("invalid monitor url")
	ErrInvalidMonitorInterval = errors.New("invalid monitor interval")
)

// Handler
var (
	ErrSearchParameters = errors.New("no search parameters")
)

func New(message string) error {
	return errors.New(message)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
