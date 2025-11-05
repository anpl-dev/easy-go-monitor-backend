package controller

import (
	"context"
	"easy-go-monitor/internal/user/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockDeleteUserUC struct {
	err error
}

func (m *mockDeleteUserUC) Execute(_ context.Context, _ usecase.DeleteUserInput) error {
	return m.err
}

func TestDeleteUserController_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		targetID       string
		ucMock         usecase.DeleteUserUseCase
		wantStatusCode int
	}{
		{
			name:           "success: delete user",
			targetID:       "11111111-1111-1111-1111-111111111111",
			ucMock:         &mockDeleteUserUC{err: nil},
			wantStatusCode: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			h := NewDeleteUserController(tt.ucMock)
			r.DELETE("/users/:id", h.Handle)

			req := httptest.NewRequest(http.MethodDelete, "/users/"+tt.targetID, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatusCode, w.Code)
		})
	}
}
