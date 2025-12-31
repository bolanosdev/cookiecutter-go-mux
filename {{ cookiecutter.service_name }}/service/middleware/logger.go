package middleware

import (
	"context"
	"net/http"
	"time"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/config"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/consts"

	"github.com/bolanosdev/go-snacks/collections"
	"github.com/bolanosdev/go-snacks/observability/logging"
	"go.opentelemetry.io/otel/trace"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (m *Middleware) Logging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			trace_id := trace.SpanContextFromContext(r.Context()).TraceID()
			logger := logging.NewContextLogger(trace_id.String(), m.Config.SERVICE.MODE.String())

			start_time := time.Now()
			method := r.Method
			url := r.RequestURI

			// Wrap the response writer to capture status code
			rw := &responseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK, // default to 200 if WriteHeader is never called
			}

			ctx := context.WithValue(r.Context(), consts.LoggerKey, &logger)
			next.ServeHTTP(rw, r.WithContext(ctx))

			duration := time.Since(start_time)
			is_ignored := is_ignored_req(r, m.Config)

			if !is_ignored {
				logger.
					Info().
					Str("trace_id", trace_id.String()).
					Str("method", method).
					Str("url", url).
					Dur("duration", duration).
					Int("status_code", rw.statusCode).
					Msg("http_logger")
			}
		})
	}
}

func is_ignored_req(r *http.Request, config config.AppConfig) bool {
	ignored_paths := collections.List[string](config.OBSERVABILITY.IGNORED_PATHS)
	_, ok := ignored_paths.Find(r.RequestURI)
	return ok
}
