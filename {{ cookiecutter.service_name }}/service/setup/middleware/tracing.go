package middleware

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

func (m *Middleware) Tracing(h func(w http.ResponseWriter, r *http.Request), operation_name string) http.Handler {
	tp := otel.GetTracerProvider()
	handler := otelhttp.NewHandler(
		http.HandlerFunc(h),
		operation_name,
		otelhttp.WithTracerProvider(tp),
	)

	return handler
}
