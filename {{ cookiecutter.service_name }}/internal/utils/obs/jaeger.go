package obs

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/config"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/consts"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TracerInterface defines the interface for tracing operations
type TracerInterface interface {
	Trace(c context.Context, name string) (context.Context, trace.Span)
	TraceFunc(c context.Context) context.Context
	TraceDB(c context.Context, query string, args interface{}) context.Context
}

type Tracer struct {
	name string
	ctx  context.Context
	tp   trace.TracerProvider
	cfg  config.ObsConfig
}

func NewTracer(ctx context.Context, service_name string, cfg config.ObsConfig) Tracer {
	tp := otel.GetTracerProvider()
	return Tracer{
		name: service_name,
		ctx:  ctx,
		tp:   tp,
		cfg:  cfg,
	}
}

func (t Tracer) Initialize() (Tracer, error) {
	// dont register trace provider if JAEGER information isnt provided through app.yaml
	if t.cfg.JAEGER.DIAL_HOSTNAME == "" {
		return t, errors.New("missing jaeger dial hostname")
	}

	res, err := resource.New(t.ctx, resource.WithAttributes(
		semconv.ServiceName(t.name),
	))

	if err != nil {
		return t, errors.Wrap(err, "failed to create resource for jaeger")
	}

	ctx, cancel := context.WithTimeout(t.ctx, time.Second)
	defer cancel()

	conn, err := grpc.NewClient(t.cfg.JAEGER.DIAL_HOSTNAME,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return t, errors.Wrap(err, "failed to create grpc connection for jaeger")
	}

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn), otlptracegrpc.WithTimeout(1000*time.Millisecond))
	if err != nil {
		return t, errors.Wrap(err, "failed to create exporter for jaeger")
	}

	processor := sdktrace.NewSimpleSpanProcessor(exporter)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(processor),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return t, nil
}

func (t Tracer) TraceFunc(ctx context.Context) context.Context {
	pc, _, _, _ := runtime.Caller(1)

	funcName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(funcName, ".")
	spanName := parts[len(parts)-1]
	if len(parts) >= 2 {
		receiver := parts[len(parts)-2]
		receiver = strings.TrimPrefix(receiver, "(*")
		receiver = strings.TrimSuffix(receiver, ")")
		spanName = receiver + "." + spanName
	}

	tracedCtx, span := t.Trace(ctx, spanName)
	defer span.End()

	return tracedCtx
}

func (t Tracer) TraceDB(ctx context.Context, query string, args interface{}) context.Context {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(funcName, ".")
	spanName := parts[len(parts)-1]
	if len(parts) >= 2 {
		receiver := parts[len(parts)-2]
		receiver = strings.TrimPrefix(receiver, "(*")
		receiver = strings.TrimSuffix(receiver, ")")
		spanName = receiver + "." + spanName
	}

	tracedCtx, span := t.Trace(ctx, spanName)
	defer span.End()

	span.SetAttributes(
		attribute.String("db.statement", query),
	)

	if args != nil {
		argsStr := fmt.Sprintf("%+v", args)
		maskedArgs := MaskSensitiveData(argsStr)
		span.SetAttributes(
			attribute.String("db.args", maskedArgs),
		)
	}

	return tracedCtx
}

func (t Tracer) Trace(c context.Context, name string) (context.Context, trace.Span) {
	ctx, span := t.tp.Tracer(t.name).Start(c, name)

	return ctx, span
}

func MaskSensitiveData(argsStr string) string {
	masked := argsStr

	for _, keyword := range consts.SensitiveKeywords {
		// Match patterns like: keyword:value or keyword: value
		// This handles both struct format and map format
		if strings.Contains(strings.ToLower(masked), strings.ToLower(keyword)) {
			// Find the keyword and mask the value after it
			lowerMasked := strings.ToLower(masked)
			idx := strings.Index(lowerMasked, strings.ToLower(keyword))

			if idx != -1 {
				// Find the start of the value (after : or :space)
				valueStart := idx + len(keyword)
				for valueStart < len(masked) && (masked[valueStart] == ':' || masked[valueStart] == ' ') {
					valueStart++
				}

				// Find the end of the value (space, comma, or closing bracket)
				valueEnd := valueStart
				for valueEnd < len(masked) && masked[valueEnd] != ' ' && masked[valueEnd] != ',' && masked[valueEnd] != '}' && masked[valueEnd] != ']' {
					valueEnd++
				}

				// Replace the value with asterisks
				if valueEnd > valueStart {
					masked = masked[:valueStart] + "***" + masked[valueEnd:]
				}
			}
		}
	}

	return masked
}
