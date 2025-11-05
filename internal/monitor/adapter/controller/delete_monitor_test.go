package controller

import (
	"context"
	"easy-go-monitor/internal/monitor/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockDeleteMonitorUC struct {
	err error
}

func (m *mockDeleteMonitorUC) Execute(_ context.Context, _ usecase.DeleteMonitorInput) error {
	return m.err
}

func TestDeleteMonitorController_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		targetID       string
		ucMock         usecase.DeleteMonitorUseCase
		wantStatusCode int
	}{
		{
			name:           "success: delete monitor",
			targetID:       "11111111-1111-1111-1111-111111111111",
			ucMock:         &mockDeleteMonitorUC{err: nil},
			wantStatusCode: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			h := NewDeleteMonitorController(tt.ucMock)
			r.DELETE("/monitors/:id", h.Handle)

			req := httptest.NewRequest(http.MethodDelete, "/monitors/"+tt.targetID, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatusCode, w.Code)
		})
	}
}
