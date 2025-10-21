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

type mockCreateUserUC struct {
	result usecase.CreateUserOutput
	err    error
}

func (m *mockCreateUserUC) Execute(_ context.Context, _ usecase.CreateUserInput) (usecase.CreateUserOutput, error) {
	return m.result, m.err
}

func TestCreateUserController_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	payload, _ := json.Marshal(map[string]string{
		"name":     "Alice",
		"email":    "alice@example.com",
		"password": "plainPassword",
	})
	wantOutput := usecase.CreateUserOutput{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		Email:     "alice@example.com",
		CreatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
	}
	wantBody := map[string]interface{}{
		"code":    float64(201),
		"message": "success",
		"data": map[string]interface{}{
			"id":         "11111111-1111-1111-1111-111111111111",
			"name":       "Alice",
			"email":      "alice@example.com",
			"created_at": "2025-04-01T00:00:00+09:00",
			"updated_at": "2025-04-01T00:00:00+09:00",
		},
	}

	tests := []struct {
		name           string
		rawPayload     []byte
		ucMock         usecase.CreateUserUseCase
		wantStatusCode int
		wantBody       map[string]interface{}
	}{
		{
			name:       "success: create user",
			rawPayload: payload,
			ucMock: &mockCreateUserUC{
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
			h := NewCreateUserController(tt.ucMock)
			r.POST("/users", h.Handle)

			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(tt.rawPayload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatusCode, w.Code)
			wantJSON, _ := json.Marshal(tt.wantBody)
			assert.JSONEq(t, string(wantJSON), w.Body.String())
		})
	}
}
