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

type mockUpdateUserUC struct {
	result usecase.UpdateUserOutput
	err    error
}

func (m *mockUpdateUserUC) Execute(_ context.Context, _ usecase.UpdateUserInput) (usecase.UpdateUserOutput, error) {
	return m.result, m.err
}

func TestUpdateUserHandler_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	targetID := "11111111-1111-1111-1111-111111111111"

	payload, _ := json.Marshal(map[string]interface{}{
		"name":          "Alice",
		"email":         "alice@example.com",
		"password_hash": "hashedPass",
	})

	want := usecase.UpdateUserOutput{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		Email:     "alice@example.com",
		UpdatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
	}

	tests := []struct {
		name           string
		rawPayload     []byte
		ucMock         usecase.UpdateUserUseCase
		wantStatusCode int
		wantBody       usecase.UpdateUserOutput
	}{
		{
			name:       "success: update monitor",
			rawPayload: payload,
			ucMock: &mockUpdateUserUC{
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
			h := NewUpdateUserHandler(tt.ucMock)
			r.PUT("/users/:id", h.Handle)

			req := httptest.NewRequest(http.MethodPut, "/users/"+targetID, bytes.NewBuffer(tt.rawPayload))
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
