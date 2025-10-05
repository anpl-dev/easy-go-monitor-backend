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

type mockCreateMonitorUC struct {
	result usecase.CreateMonitorOutput
	err    error
}

func (m *mockCreateMonitorUC) Execute(_ context.Context, _ usecase.CreateMonitorInput) (usecase.CreateMonitorOutput, error) {
	return m.result, m.err
}

func TestCreateMonitorHandler_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	payload, _ := json.Marshal(map[string]interface{}{
		"user_id":         "11111111-1111-1111-1111-111111111111",
		"name":            "test-monitor",
		"url":             "https://example.com",
		"interval_second": 60,
	})

	want := usecase.CreateMonitorOutput{
		ID:             uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:           "Alice",
		URL:            "https://example.com",
		IntervalSecond: 60,
		CreatedAt:      time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt:      time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
	}

	tests := []struct {
		name           string
		rawPayload     []byte
		ucMock         usecase.CreateMonitorUseCase
		wantStatusCode int
		wantBody       usecase.CreateMonitorOutput
	}{
		{
			name:       "success: create user",
			rawPayload: payload,
			ucMock: &mockCreateMonitorUC{
				result: want,
				err:    nil,
			},
			wantStatusCode: http.StatusCreated,
			wantBody:       want,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			h := NewCreateMonitorHandler(tt.ucMock)
			r.POST("/monitors", h.Handle)

			req := httptest.NewRequest(http.MethodPost, "/monitors", bytes.NewBuffer(tt.rawPayload))
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
