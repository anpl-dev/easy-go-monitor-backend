package controller

import (
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

type mockFindMonitorByIDUC struct {
	result usecase.FindMonitorByIDOutput
	err    error
}

func (m *mockFindMonitorByIDUC) Execute(_ context.Context, _ usecase.FindMonitorByIDInput) (usecase.FindMonitorByIDOutput, error) {
	return m.result, m.err
}

func TestFindMonitorByIDController_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	wantOUtput := usecase.FindMonitorByIDOutput{
		ID:             uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:           "test-monitor",
		URL:            "https://example.com",
		CreatedAt:      time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
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
			"created_at":      "2025-04-01T00:00:00+09:00",
			"updated_at":      "2025-04-01T00:00:00+09:00",
		},
	}

	tests := []struct {
		name           string
		targetID       string
		ucMock         usecase.FindMonitorByIDUseCase
		wantStatusCode int
		wantBody       map[string]interface{}
	}{
		{
			name:     "success: find monitor by id",
			targetID: "11111111-1111-1111-1111-111111111111",
			ucMock: &mockFindMonitorByIDUC{
				result: wantOUtput,
				err:    nil,
			},
			wantStatusCode: http.StatusOK,
			wantBody:       wantBody,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			h := NewFindMonitorByIDController(tt.ucMock)
			r.GET("/monitors/:id", h.Handle)

			req := httptest.NewRequest(http.MethodGet, "/monitors/"+tt.targetID, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatusCode, w.Code)
			wantJSON, _ := json.Marshal(tt.wantBody)
			assert.JSONEq(t, string(wantJSON), w.Body.String())
		})
	}
}
