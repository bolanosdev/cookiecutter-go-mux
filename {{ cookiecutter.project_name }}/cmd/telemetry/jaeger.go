package telemetry

import (
	"context"
	"errors"
	"time"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/cfg"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ErrFailedTracerSetup = errors.New("failed to setup new tracer")

func Initialize(ctx context.Context, config cfg.AppConfig) error {
	// dont register trace provider if JAEGER information isnt provided through app.yaml
	if config.OBSERVABILITY.JAEGER.DIAL_HOSTNAME == "" {
		return nil
	}

	res, err := resource.New(ctx, resource.WithAttributes(
		semconv.ServiceName(config.SERVICE_NAME),
	))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	conn, err := grpc.NewClient(config.OBSERVABILITY.JAEGER.DIAL_HOSTNAME,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn), otlptracegrpc.WithTimeout(1000*time.Millisecond))
	if err != nil {
		return err
	}

	processor := sdktrace.NewSimpleSpanProcessor(exporter)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(processor),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return nil
}
