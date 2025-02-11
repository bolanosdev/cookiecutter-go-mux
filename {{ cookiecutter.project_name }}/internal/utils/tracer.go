package utils

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/cfg"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func TracerWithContext(c context.Context, name string) (context.Context, trace.Span) {
	config := cfg.Load(".")
	tp := otel.GetTracerProvider()
	ctx, span := tp.Tracer(config.SERVICE_NAME).Start(c, name)

	return ctx, span
}
