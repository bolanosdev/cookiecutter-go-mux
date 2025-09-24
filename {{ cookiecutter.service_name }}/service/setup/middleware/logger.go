package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/config"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/utils"

	"github.com/rs/zerolog"
)

func (m *Middleware) Logging(logger zerolog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start_time := time.Now()
			method := r.Method
			url := r.RequestURI
			req_body := get_request_body(r, m.Config)

			next.ServeHTTP(w, r)

			status_code := w.Header().Get("x-status-code")
			trace_id := w.Header().Get("x-trace-id")
			message := w.Header().Get("x-error-message")

			duration := time.Since(start_time)
			is_ignored := is_ignored_req(r, m.Config)

			if !is_ignored {
				logger.Info().
					Str("trace_id", trace_id).
					Str("method", method).
					Str("url", url).
					Dur("duration", duration).
					Bytes("request", req_body).
					Str("status_code", status_code).
					Str("message", message).
					Msg("http_logger")
			}
		})
	}
}

func is_ignored_req(r *http.Request, config config.AppConfig) bool {
	return utils.Contains(config.OBSERVABILITY.IGNORED_PATHS, func(path string) bool { return path == r.RequestURI })
}

func is_sensitive_req(r *http.Request, config config.AppConfig) bool {
	return utils.Contains(config.OBSERVABILITY.SENSITIVE_PATHS, func(path string) bool { return path == r.RequestURI })
}

func get_request_body(r *http.Request, config config.AppConfig) []byte {
	if !is_sensitive_req(r, config) {
		var buf bytes.Buffer
		tee := io.TeeReader(r.Body, &buf)
		body, _ := io.ReadAll(tee)
		r.Body = io.NopCloser(&buf)
		// result := string(body)
		return body
	} else {
		return []byte{}
	}
}
