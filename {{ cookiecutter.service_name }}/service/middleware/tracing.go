package middleware

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

func (m *Middleware) Tracing() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip tracing for ignored paths
			for _, ignoredPath := range m.Config.OBSERVABILITY.IGNORED_PATHS {
				if r.URL.Path == ignoredPath {
					next.ServeHTTP(w, r)
					return
				}
			}

			// Extract method and path for tracer label
			method := r.Method
			path := r.URL.Path
			tracerLabel := method + " " + path

			tp := otel.GetTracerProvider()

			handler := otelhttp.NewHandler(
				next,
				tracerLabel,
				otelhttp.WithTracerProvider(tp),
			)

			handler.ServeHTTP(w, r)
		})
	}
}
