package handler

import (
	"context"
	"encoding/json"
	"go-monitor-tool/internal/usecase"
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

func TestFindMonitorByIDHandler_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	want := usecase.FindMonitorByIDOutput{
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
		targetID       string
		ucMock         usecase.FindMonitorByIDUseCase
		wantStatusCode int
		wantBody       usecase.FindMonitorByIDOutput
	}{
		{
			name:     "success: find monitor by id handler",
			targetID: "11111111-1111-1111-1111-111111111111",
			ucMock: &mockFindMonitorByIDUC{
				result: usecase.FindMonitorByIDOutput{
					ID:             uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					UserID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Name:           "Alice",
					URL:            "https://example.com",
					IntervalSecond: 60,
					CreatedAt:      time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
					UpdatedAt:      time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
				},
				err: nil,
			},
			wantStatusCode: http.StatusOK,
			wantBody:       want,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			h := NewFindMonitorByIDHandler(tt.ucMock)
			r.GET("/monitors/:id", h.Handle)

			req := httptest.NewRequest(http.MethodGet, "/monitors/"+tt.targetID, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatusCode {
				t.Errorf("[%s] status code: got %v, want %v", tt.name, w.Code, tt.wantStatusCode)
			}
			wantJSON, _ := json.Marshal(tt.wantBody)
			assert.JSONEq(t, string(wantJSON), w.Body.String())
		})
	}
}
