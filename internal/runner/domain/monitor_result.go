package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	MonitorResult struct {
		MonitorID uuid.UUID
		Status    string
		LatencyMs int64
		Timestamp time.Time
	}
)
