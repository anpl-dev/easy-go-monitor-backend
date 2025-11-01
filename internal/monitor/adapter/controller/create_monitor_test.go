package controller

import (
	"bytes"
	"context"
	"easy-go-monitor/internal/infra/logger"
	"easy-go-monitor/internal/monitor/usecase"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockCreateMonitorUC struct {
	result usecase.CreateMonitorOutput
	err    error
}

func (m *mockCreateMonitorUC) Execute(_ context.Context, _ usecase.CreateMonitorInput) (usecase.CreateMonitorOutput, error) {
	return m.result, m.err
}

func TestCreateMonitorController_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	payload, _ := json.Marshal(map[string]interface{}{
		"user_id":         "11111111-1111-1111-1111-111111111111",
		"name":            "test-monitor",
		"url":             "https://example.com",
		"interval_second": 60,
	})
	wantOutput := usecase.CreateMonitorOutput{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:    uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "test-monitor",
		URL:       "https://example.com",
		CreatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
	}
	wantBody := map[string]interface{}{
		"code":    float64(201),
		"message": "success",
		"data": map[string]interface{}{
			"id":              "11111111-1111-1111-1111-111111111111",
			"user_id":         "11111111-1111-1111-1111-111111111111",
			"name":            "test-monitor",
			"url":             "https://example.com",
			"interval_second": float64(60),
			"created_at":      "2025-04-01T00:00:00+09:00",
			"updated_at":      "2025-04-01T00:00:00+09:00",
		},
	}

	tests := []struct {
		name           string
		rawPayload     []byte
		ucMock         usecase.CreateMonitorUseCase
		wantStatusCode int
		wantBody       map[string]interface{}
	}{
		{
			name:       "success: create user",
			rawPayload: payload,
			ucMock: &mockCreateMonitorUC{
				result: wantOutput,
				err:    nil,
			},
			wantStatusCode: http.StatusCreated,
			wantBody:       wantBody,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			h := NewCreateMonitorController(tt.ucMock, &logger.Logger{})
			r.POST("/monitors", h.Handle)

			req := httptest.NewRequest(http.MethodPost, "/monitors", bytes.NewBuffer(tt.rawPayload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatusCode, w.Code)
			wantJSON, _ := json.Marshal(tt.wantBody)
			assert.JSONEq(t, string(wantJSON), w.Body.String())
		})
	}
}
