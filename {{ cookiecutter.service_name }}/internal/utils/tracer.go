package utils

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/config"

	"go.opentelemetry.io/otel/trace"
)

type Tracer struct {
	name string
	tp   trace.TracerProvider
	cfg  config.AppConfig
}

func NewTracer(tp trace.TracerProvider, cfg config.AppConfig) Tracer {
	return Tracer{
		tp:  tp,
		cfg: cfg,
	}
}

func (t Tracer) Trace(c context.Context, name string) (context.Context, trace.Span) {
	ctx, span := t.tp.Tracer(t.cfg.SERVICE.NAME).Start(c, name)

	return ctx, span
}
