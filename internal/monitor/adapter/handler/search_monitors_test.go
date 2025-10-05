package handler

import (
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

type mockSearchMonitorsUC struct {
	result []usecase.SearchMonitorsOutput
	err    error
}

func (m *mockSearchMonitorsUC) Execute(_ context.Context, _ usecase.SearchMonitorsInput) ([]usecase.SearchMonitorsOutput, error) {
	return m.result, m.err
}

func TestSearchMonitorsHandler_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	user := usecase.SearchMonitorsOutput{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		UserID:    uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		URL:       "https://example.com",
		CreatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
	}

	tests := []struct {
		name           string
		queryParam     string
		ucMock         usecase.SearchMonitorsUseCase
		wantStatusCode int
		wantBody       []usecase.SearchMonitorsOutput
	}{
		{
			name:       "success: search monitors",
			queryParam: "?user_id=11111111-1111-1111-1111-111111111111",
			ucMock: &mockSearchMonitorsUC{
				result: []usecase.SearchMonitorsOutput{user},
				err:    nil,
			},
			wantStatusCode: http.StatusOK,
			wantBody:       []usecase.SearchMonitorsOutput{user},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			h := NewSearchMonitorsHandler(tt.ucMock)
			r.GET("/monitors/search", h.Handle)

			req := httptest.NewRequest(http.MethodGet, "/monitors/search"+tt.queryParam, nil)
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
