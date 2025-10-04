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

type mockFindUserByIDUC struct {
	result usecase.FindUserByIDOutput
	err    error
}

func (m *mockFindUserByIDUC) Execute(_ context.Context, _ usecase.FindUserByIDInput) (usecase.FindUserByIDOutput, error) {
	return m.result, m.err
}

func TestFindUserByIDHandler_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	want := usecase.FindUserByIDOutput{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		Email:     "alice@example.com",
		CreatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
	}

	tests := []struct {
		name           string
		targetID       string
		ucMock         usecase.FindUserByIDUseCase
		wantStatusCode int
		wantBody       usecase.FindUserByIDOutput
	}{
		{
			name:     "success: find user by id",
			targetID: "11111111-1111-1111-1111-111111111111",
			ucMock: &mockFindUserByIDUC{
				result: usecase.FindUserByIDOutput{
					ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					Name:      "Alice",
					Email:     "alice@example.com",
					CreatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
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
			h := NewFindUserByIDHandler(tt.ucMock)
			r.GET("/users/:id", h.Handle)

			req := httptest.NewRequest(http.MethodGet, "/users/"+tt.targetID, nil)
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
