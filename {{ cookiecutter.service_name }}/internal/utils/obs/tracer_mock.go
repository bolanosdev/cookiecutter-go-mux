package obs

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/config"

	"go.opentelemetry.io/otel/trace"
)

type MockTracer struct {
	name string
	cfg  config.AppConfig
}

func NewMockTracer(cfg config.AppConfig) MockTracer {
	return MockTracer{
		cfg: cfg,
	}
}

func (m MockTracer) Trace(c context.Context, name string) (context.Context, trace.Span) {
	return c, trace.SpanFromContext(c)
}
