package utils

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/cfg"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func TracerWithGinContext(c *gin.Context, name string) (context.Context, trace.Span) {
	config := cfg.Load(".")
	tp := otel.GetTracerProvider()

	ctx, span := tp.Tracer(config.SERVICE_NAME).Start(c, name)
	trace_id := trace.SpanContextFromContext(ctx).TraceID().String()

	c.Set("x-trace-id", trace_id)
	return ctx, span
}

func TracerWithContext(c context.Context, name string) (context.Context, trace.Span) {
	config := cfg.Load(".")
	tp := otel.GetTracerProvider()
	ctx, span := tp.Tracer(config.SERVICE_NAME).Start(c, name)

	return ctx, span
}
