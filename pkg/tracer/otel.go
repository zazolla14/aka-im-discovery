package tracer

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/1nterdigital/aka-im-discover/pkg/common/config"
)

type OtelTracer struct {
	ServiceName attribute.KeyValue
	TracerConf  *config.Tracer
}

func NewOtelTracer(tracerConf *config.Tracer) (otelTracer *OtelTracer) {
	return &OtelTracer{
		ServiceName: semconv.ServiceNameKey.String(tracerConf.AppName.Api),
		TracerConf:  tracerConf,
	}
}

// InitTracer Initialize a gRPC connection to be used by both the tracer providers.
func (ot *OtelTracer) InitTracer(ctx context.Context) (func(context.Context) error, error) {
	conn, err := grpc.NewClient(
		ot.TracerConf.Otel.Collector.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection: %w", err)
	}

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(ot.TracerConf.AppName.Api),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// TracerProvider with BatchSpanProcessor
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(traceExporter),
	)
	otel.SetTracerProvider(tp)

	// Context propagation (trace headers between services)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tp.Shutdown, nil
}
