package utils

import (
	"context"
	"log"
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/cfg"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func Instrument(h func(w http.ResponseWriter, r *http.Request), operation_name string) http.Handler {
	tp := otel.GetTracerProvider()

	handler := otelhttp.NewHandler(
		http.HandlerFunc(h),
		operation_name,
		otelhttp.WithTracerProvider(tp),
	)
	return handler
}

func TracerWithGinContext(c context.Context, name string) (context.Context, trace.Span) {
	config := cfg.Load(".")
	tp := otel.GetTracerProvider()
	ctx, span := tp.Tracer(config.SERVICE_NAME).Start(
		c,
		name,
	)

	trace_id := trace.SpanContextFromContext(ctx).TraceID().String()
	log.Printf("trace_id $s", trace_id)

	// c.Set("x-trace-id", trace_id)
	return ctx, span
}

func TracerWithContext(c context.Context, name string) (context.Context, trace.Span) {
	config := cfg.Load(".")
	tp := otel.GetTracerProvider()
	ctx, span := tp.Tracer(config.SERVICE_NAME).Start(c, name)

	return ctx, span
}
