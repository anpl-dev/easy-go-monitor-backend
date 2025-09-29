package handler

import (
	"bytes"
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
		"name":          "Alice",
		"email":         "alice@example.com",
		"password_hash": "hashedPass",
	})

	expected := usecase.CreateUserOutput{
		ID:           uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:         "Alice",
		Email:        "alice@example.com",
		PasswordHash: "hashedPass",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	tests := []struct {
		name               string
		rawPayload         []byte
		ucMock             usecase.CreateUserUseCase
		expectedStatusCode int
		expectedBody       usecase.CreateUserOutput
	}{
		{
			name:       "success: create user handler",
			rawPayload: payload,
			ucMock: &mockCreateUserUC{
				result: usecase.CreateUserOutput{
					ID:           uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Name:         "Alice",
					Email:        "alice@example.com",
					PasswordHash: "hashedPass",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				},
				err: nil,
			},
			expectedStatusCode: http.StatusCreated,
			expectedBody:       expected,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(tt.rawPayload))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			h := NewCreateUserHandler(tt.ucMock)
			h.Handle(c)

			if w.Code != tt.expectedStatusCode {
				t.Errorf("[%s] status code: got %v, want %v", tt.name, w.Code, tt.expectedStatusCode)
			}

			expectedJSON, _ := json.Marshal(expected)

			assert.JSONEq(t, string(expectedJSON), w.Body.String())
		})
	}
}
