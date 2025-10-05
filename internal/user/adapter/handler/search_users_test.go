package handler

import (
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

type mockSearchUsersUC struct {
	result []usecase.SearchUsersOutput
	err    error
}

func (m *mockSearchUsersUC) Execute(_ context.Context, _ usecase.SearchUsersInput) ([]usecase.SearchUsersOutput, error) {
	return m.result, m.err
}

func TestSearchUsersHandler_Execute(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	user := usecase.SearchUsersOutput{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Name:      "Alice",
		Email:     "alice@example.com",
		CreatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local),
	}

	tests := []struct {
		name           string
		queryParam     string
		ucMock         usecase.SearchUsersUseCase
		wantStatusCode int
		wantBody       []usecase.SearchUsersOutput
	}{
		{
			name:       "success: search users",
			queryParam: "?name=Alice",
			ucMock: &mockSearchUsersUC{
				result: []usecase.SearchUsersOutput{user},
				err:    nil,
			},
			wantStatusCode: http.StatusOK,
			wantBody:       []usecase.SearchUsersOutput{user},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			h := NewSearchUsersHandler(tt.ucMock)
			r.GET("/users/search/", h.Handle)

			req := httptest.NewRequest(http.MethodGet, "/users/search/"+tt.queryParam, nil)
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
