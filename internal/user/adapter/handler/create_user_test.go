package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"go-monitor-tool/internal/user/usecase"
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

func TestCreateUserHandler_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	payload, _ := json.Marshal(map[string]string{
		"name":     "Alice",
		"email":    "alice@example.com",
		"password": "plainPassword",
	})

	want := usecase.CreateUserOutput{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		Email:     "alice@example.com",
		CreatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
	}

	tests := []struct {
		name           string
		rawPayload     []byte
		ucMock         usecase.CreateUserUseCase
		wantStatusCode int
		wantBody       usecase.CreateUserOutput
	}{
		{
			name:       "success: create user",
			rawPayload: payload,
			ucMock: &mockCreateUserUC{
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
			h := NewCreateUserHandler(tt.ucMock)
			r.POST("/users", h.Handle)

			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(tt.rawPayload))
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
