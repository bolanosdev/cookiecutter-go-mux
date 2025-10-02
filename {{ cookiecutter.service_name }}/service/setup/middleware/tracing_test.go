package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTracingMiddleware(t *testing.T) {
	handler := m.Tracing(mockHandler, "test-op")
	rr := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	handler.ServeHTTP(rr, req)
	require.Equal(t, rr.Result().StatusCode, 200)
}
