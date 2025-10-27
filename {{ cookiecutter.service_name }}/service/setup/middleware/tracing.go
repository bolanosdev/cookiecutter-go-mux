package middleware

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

func (m *Middleware) Tracing() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tp := otel.GetTracerProvider()

			handler := otelhttp.NewHandler(
				next,
				"http_tracer",
				otelhttp.WithTracerProvider(tp),
			)

			handler.ServeHTTP(w, r)
		})
	}
}
