package controller

import (
	"bytes"
	"context"
	"easy-go-monitor/internal/user/usecase"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockUpdateUserUC struct {
	result usecase.UpdateUserOutput
	err    error
}

func (m *mockUpdateUserUC) Execute(_ context.Context, _ usecase.UpdateUserInput) (usecase.UpdateUserOutput, error) {
	return m.result, m.err
}

func TestUpdateUserController_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	targetID := "11111111-1111-1111-1111-111111111111"
	payload, _ := json.Marshal(map[string]interface{}{
		"name":     "Alice",
		"email":    "alice@example.com",
		"password": "hashedPass",
	})
	wantOutput := usecase.UpdateUserOutput{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		Email:     "alice@example.com",
		UpdatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
	}
	wantBody := map[string]interface{}{
		"code":    float64(200),
		"message": "success",
		"data": map[string]interface{}{
			"id":         "11111111-1111-1111-1111-111111111111",
			"name":       "Alice",
			"email":      "alice@example.com",
			"updated_at": "2025-04-01T00:00:00+09:00",
		},
	}

	tests := []struct {
		name           string
		rawPayload     []byte
		ucMock         usecase.UpdateUserUseCase
		wantStatusCode int
		wantBody       map[string]interface{}
	}{
		{
			name:       "success: update monitor",
			rawPayload: payload,
			ucMock: &mockUpdateUserUC{
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
			h := NewUpdateUserController(tt.ucMock)
			r.PUT("/users/:id", h.Handle)

			req := httptest.NewRequest(http.MethodPut, "/users/"+targetID, bytes.NewBuffer(tt.rawPayload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatusCode, w.Code)
			wantJSON, _ := json.Marshal(tt.wantBody)
			assert.JSONEq(t, string(wantJSON), w.Body.String())
		})
	}
}
