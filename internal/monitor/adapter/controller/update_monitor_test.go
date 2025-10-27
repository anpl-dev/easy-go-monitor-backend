package controller

import (
	"bytes"
	"context"
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

type mockUpdateMonitorUC struct {
	result usecase.UpdateMonitorOutput
	err    error
}

func (m *mockUpdateMonitorUC) Execute(_ context.Context, _ usecase.UpdateMonitorInput) (usecase.UpdateMonitorOutput, error) {
	return m.result, m.err
}

func TestUpdateMonitorController_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	targetID := "11111111-1111-1111-1111-111111111111"
	payload, _ := json.Marshal(map[string]interface{}{
		"name":            "test-monitor",
		"url":             "https://example.com",
		"interval_second": 60,
	})
	wantOutput := usecase.UpdateMonitorOutput{
		ID:             uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:           "test-monitor",
		URL:            "https://example.com",
		UpdatedAt:      time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
	}
	wantBody := map[string]interface{}{
		"code":    float64(200),
		"message": "success",
		"data": map[string]interface{}{
			"id":              "11111111-1111-1111-1111-111111111111",
			"user_id":         "11111111-1111-1111-1111-111111111111",
			"name":            "test-monitor",
			"url":             "https://example.com",
			"interval_second": float64(60),
			"updated_at":      "2025-04-01T00:00:00+09:00",
		},
	}

	tests := []struct {
		name           string
		rawPayload     []byte
		ucMock         usecase.UpdateMonitorUseCase
		wantStatusCode int
		wantBody       map[string]interface{}
	}{
		{
			name:       "success: update monitor",
			rawPayload: payload,
			ucMock: &mockUpdateMonitorUC{
				result: wantOutput,
				err:    nil,
			},
			wantStatusCode: http.StatusOK,
			wantBody:       wantBody,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			h := NewUpdateMonitorController(tt.ucMock)
			r.PUT("/monitors/:id", h.Handle)

			req := httptest.NewRequest(http.MethodPut, "/monitors/"+targetID, bytes.NewBuffer(tt.rawPayload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatusCode, w.Code)
			wantJSON, _ := json.Marshal(tt.wantBody)
			assert.JSONEq(t, string(wantJSON), w.Body.String())
		})
	}
}
