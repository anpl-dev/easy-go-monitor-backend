package domain

import "github.com/google/uuid"

type (
	Monitor struct {
		ID   uuid.UUID
		URL  string
		Type string
	}
)
