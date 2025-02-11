package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/cfg"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/utils"

	"github.com/rs/zerolog/log"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config := cfg.Load(".")

		start_time := time.Now()
		method := r.Method
		url := r.RequestURI
		req_body := get_request_body(r, config)

		next.ServeHTTP(w, r)

		status_code := w.Header().Get("x-status-code")
		trace_id := w.Header().Get("x-trace-id")
		message := w.Header().Get("x-error-message")
		duration := time.Since(start_time)
		is_ignored := is_ignored_req(r, config)

		if !is_ignored {
			log.Info().
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

func is_ignored_req(r *http.Request, config cfg.AppConfig) bool {
	return utils.Contains(config.OBSERVABILITY.IGNORED_PATHS, func(path string) bool { return path == r.RequestURI })
}

func is_sensitive_req(r *http.Request, config cfg.AppConfig) bool {
	return utils.Contains(config.OBSERVABILITY.SENSITIVE_PATHS, func(path string) bool { return path == r.RequestURI })
}

func get_request_body(r *http.Request, config cfg.AppConfig) []byte {
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
