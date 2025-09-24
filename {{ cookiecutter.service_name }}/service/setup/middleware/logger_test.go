package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestLoggerMiddleware(t *testing.T) {
	var buf bytes.Buffer
	logger := zerolog.New(&buf)
	handler := m.Logging(logger)(mockSuccessHandlerFunc)
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(rr, req)

	string_buff := buf.String()
	require.Equal(t, rr.Result().StatusCode, 200)

	require.Contains(t, string_buff, `"http_logger"`)
	require.Contains(t, string_buff, `"level":"info"`)
	require.Contains(t, string_buff, `"method":"GET"`)
	require.Contains(t, string_buff, `"url":"/test"`)
	require.Contains(t, string_buff, `"message":""`)
	require.Contains(t, string_buff, `"request":""`)
	require.Contains(t, string_buff, `"status_code":"200"`)
	require.Contains(t, string_buff, `"trace_id":"2ec68d14-d2e7-4e1a-a21d-8615282bbac7"`)
}

func TestLoggerMiddlewareWithBody(t *testing.T) {
	var buf bytes.Buffer
	logger := zerolog.New(&buf)
	handler := m.Logging(logger)(mockSuccessHandlerFunc)

	string_body := `{"username": "test", "password": "secret"}`
	body := []byte(string_body)
	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(rr, req)

	string_buff := buf.String()
	require.Equal(t, rr.Result().StatusCode, 200)

	require.Contains(t, string_buff, `"http_logger"`)
	require.Contains(t, string_buff, `"level":"info"`)
	require.Contains(t, string_buff, `"method":"POST"`)
	require.Contains(t, string_buff, `"url":"/test"`)
	require.Contains(t, string_buff, `"message":""`)
	require.Contains(t, string_buff, `request`)
	require.Contains(t, string_buff, `username`)
	require.Contains(t, string_buff, `test`)
	require.Contains(t, string_buff, `password`)
	require.Contains(t, string_buff, `secret`)
	require.Contains(t, string_buff, `"status_code":"200"`)
	require.Contains(t, string_buff, `"trace_id":"2ec68d14-d2e7-4e1a-a21d-8615282bbac7"`)
}

func TestLoggerMiddlewareSensitivePaths(t *testing.T) {
	var buf bytes.Buffer
	logger := zerolog.New(&buf)
	handler := m.Logging(logger)(mockSuccessHandlerFunc)

	string_body := `{"username": "test", "password": "secret"}`
	body := []byte(string_body)
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(rr, req)

	string_buff := buf.String()
	require.Equal(t, rr.Result().StatusCode, 200)

	require.Contains(t, string_buff, `"http_logger"`)
	require.Contains(t, string_buff, `"level":"info"`)
	require.Contains(t, string_buff, `"method":"POST"`)
	require.Contains(t, string_buff, `"url":"/accounts"`)
	require.Contains(t, string_buff, `"message":""`)
	require.Contains(t, string_buff, `request`)
	require.NotContains(t, string_buff, `username`)
	require.NotContains(t, string_buff, `test`)
	require.NotContains(t, string_buff, `password`)
	require.NotContains(t, string_buff, `secret`)
	require.Contains(t, string_buff, `"status_code":"200"`)
	require.Contains(t, string_buff, `"trace_id":"2ec68d14-d2e7-4e1a-a21d-8615282bbac7"`)
}

func TestLoggerMiddlewareIgnoredPaths(t *testing.T) {
	var buf bytes.Buffer
	logger := zerolog.New(&buf)
	handler := m.Logging(logger)(mockSuccessHandlerFunc)

	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	rr := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(rr, req)

	string_buff := buf.String()
	require.Equal(t, rr.Result().StatusCode, 200)
	require.NotContains(t, string_buff, `"http_logger"`)
}
