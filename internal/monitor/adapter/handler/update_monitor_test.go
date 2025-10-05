package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"go-monitor-tool/internal/monitor/usecase"
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

func TestUpdateMonitorHandler_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	targetID := "11111111-1111-1111-1111-111111111111"

	payload, _ := json.Marshal(map[string]interface{}{
		"name":            "Alice",
		"url":             "https://example.com",
		"interval_second": 60,
	})

	want := usecase.UpdateMonitorOutput{
		ID:             uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:           "Alice",
		URL:            "https://example.com",
		IntervalSecond: 60,
		UpdatedAt:      time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
	}

	tests := []struct {
		name           string
		rawPayload     []byte
		ucMock         usecase.UpdateMonitorUseCase
		wantStatusCode int
		wantBody       usecase.UpdateMonitorOutput
	}{
		{
			name:       "success: update monitor",
			rawPayload: payload,
			ucMock: &mockUpdateMonitorUC{
				result: want,
				err:    nil,
			},
			wantStatusCode: http.StatusOK,
			wantBody:       want,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			h := NewUpdateMonitorHandler(tt.ucMock)
			r.PUT("/monitors/:id", h.Handle)

			req := httptest.NewRequest(http.MethodPut, "/monitors/"+targetID, bytes.NewBuffer(tt.rawPayload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatusCode {
				t.Errorf("[TestCase %s] status code: Got %v, Want %v", tt.name, w.Code, tt.wantStatusCode)
			}
			wantJSON, _ := json.Marshal(tt.wantBody)
			assert.JSONEq(t, string(wantJSON), w.Body.String())
		})
	}
}
