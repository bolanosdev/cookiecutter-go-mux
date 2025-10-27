package obs

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/config"

	"go.opentelemetry.io/otel/trace"
)

// TracerInterface defines the interface for tracing operations
type TracerInterface interface {
	Trace(c context.Context, name string) (context.Context, trace.Span)
}

// MockTracer is a mock implementation of the TracerInterface for testing
type MockTracer struct {
	name string
	cfg  config.AppConfig
}

// NewMockTracer creates a new mock tracer for testing
func NewMockTracer(cfg config.AppConfig) MockTracer {
	return MockTracer{
		cfg: cfg,
	}
}

// Trace returns a context and a no-op span for testing
func (m MockTracer) Trace(c context.Context, name string) (context.Context, trace.Span) {
	// Return the original context and a no-op span
	// This prevents any actual tracing operations during tests
	return c, trace.SpanFromContext(c)
}
